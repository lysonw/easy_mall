package script

import (
	"context"
	"easy_mall/config"
	i "easy_mall/init"
	"easy_mall/repository/cache"
	"easy_mall/repository/db/dao"
	"easy_mall/repository/db/model"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	todo:这边后续可以优化成批量写
*/

// ProductClickRecordHandleScript  注意：lua脚本会将redis的nil转换成false
var ProductClickRecordHandleScript = redis.NewScript(`
	local key = KEYS[1]
	local maxNum = tonumber(ARGV[1])

	local len = redis.call('LLEN', key)
	local index = math.min(len, maxNum)
	if index == 0 then
		return {}
	end
	
	local values = redis.call('LRANGE',key,0,index-1)
	redis.call('LTRIM',key,index,-1)
	return values
`)

var ExitChan = make(chan struct{})

func ProductClickRecordDaemon() {
	go func() {
		ctx := context.Background()
		quit := make(chan os.Signal, 3)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
		log.Println("ProductClickRecordDaemon 开始执行")

		// timeout 5s, batchNum 1000, intervalTime 1s
		go ProductClickRecordHandle(ctx, config.ProductClickRecordHandlerTimeout, config.ProductClickRecordHandlerBatchNum, config.ProductClickRecordHandlerInterval)

		<-quit
		log.Println("ProductClickRecordDaemon 执行结束")
		close(quit)
		close(ExitChan)
	}()
}

func ProductClickRecordHandle(ctx context.Context, timeout, batchNum, intervalTime int) {
loop:
	for {
		subCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		select {
		case <-ExitChan:
			log.Println("AnalysisDataHandleDaemon 收到退出信号")
			cancel()
			return
		default:
			// batchNum数量不宜过多，目前发现 batchNum = 6000时，redis（8核32G）会出现io连接超时的报错
			values, err := pullClickRecordData(subCtx, cache.ProductClickRecordCacheKey, batchNum, intervalTime)
			if err != nil {
				log.Println(err)
				goto loop
			}

			if len(values) == 0 {
				goto loop
			}

			// 这边不进行异步处理，防止数据量大的情况下开启过多的goroutine
			batchInsertClickRecordData(subCtx, values)
		}
	}
}

func pullClickRecordData(ctx context.Context, key string, batchNum, intervalTime int) ([]string, error) {
	buffer := make([]string, 0, batchNum)
loop:
	for {
		select {
		// 这边可能是超时取消，也可能是收到退出信号
		case <-ctx.Done():
			return buffer, nil
		default:
			data, err := DoProductClickRecordHandleScript(key, batchNum)
			if err != nil {
				return buffer, err
			}
			buffer = append(buffer, data...)
			if len(buffer) < batchNum {
				time.Sleep(time.Duration(intervalTime) * time.Second)
				goto loop
			}
			return buffer, nil
		}
	}
}

func DoProductClickRecordHandleScript(key string, batchNum int) (res []string, err error) {
	res, err = ProductClickRecordHandleScript.Run(context.Background(), i.RedisClient, []string{key}, batchNum).StringSlice()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func batchInsertClickRecordData(ctx context.Context, values []string) {
	records := make([]model.ClickRecord, 0, len(values))
	for _, v := range values {
		var item model.ClickRecord
		if err := json.Unmarshal([]byte(v), &item); err != nil {
			log.Println(err)
			continue
		}
		records = append(records, item)
	}

	if err := dao.NewClickRecordDao(ctx).BatchInsert(records, config.ProductClickRecordHandlerBatchNum); err != nil {
		log.Println(err)
		return
	}
}
