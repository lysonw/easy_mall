package script

import (
	"context"
	"easy_mall/repository/cache"
	"easy_mall/repository/db/dao"
	"easy_mall/repository/db/model"
	"encoding/json"
	"log"
	"time"
)

/*
	todo:这边后续可以优化成批量写
*/

func ProductClickRecordDaemon() {
	ctx := context.Background()
	c := cache.Cache{Ctx: ctx}
	for {
		res := c.BPop(cache.ProductClickRecordCacheKey, 5*time.Second)
		if len(res) > 1 {
			var record model.ClickRecord
			if err := json.Unmarshal([]byte(res[1]), &record); err != nil {
				log.Println(err)
				continue
			}

			if err := dao.NewClickRecordDao(ctx).Insert(record); err != nil {
				log.Println(err)
			}
		}
	}
}
