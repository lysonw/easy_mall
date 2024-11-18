package config

var (
	MySQLDsn string
)

var (
	RedisHost     string
	RedisPort     string
	RedisUsername string
	RedisPassword string
	RedisDbName   int
)

var (
	CategoryListShowLen = 6
)

var (
	ProductClickRecordHandlerTimeout  = 5  // 秒
	ProductClickRecordHandlerBatchNum = 50 // 条
	ProductClickRecordHandlerInterval = 1  // 秒
)
