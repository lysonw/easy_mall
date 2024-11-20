package config

import "fmt"

var (
	username = "root"      //账号
	password = "mysql"     //密码
	host     = "127.0.0.1" //数据库地址，可以是Ip或者域名
	port     = 3306        //数据库端口
	Dbname   = "easy_mall" //数据库名
	timeout  = "10s"       //连接超时，10秒
	MySQLDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
)

var (
	RedisHost     = "127.0.0.1"
	RedisPort     = "6379"
	RedisUsername = ""
	RedisPassword = ""
	RedisDbName   = 1
)

var (
	ProductClickRecordHandlerTimeout  = 5  // 秒
	ProductClickRecordHandlerBatchNum = 50 // 条
	ProductClickRecordHandlerInterval = 1  // 秒
)
