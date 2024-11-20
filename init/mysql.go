package init

import (
	"context"
	"easy_mall/config"
	"easy_mall/repository/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var DB *gorm.DB

func InitMySQL() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.MySQLDsn, // DSN data source name
		DefaultStringSize:         256,             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,           // 根据版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	DB = db

	err = db.AutoMigrate(&model.Product{}, &model.ClickRecord{})
	if err != nil {
		log.Println(err)
	}
}

func DBClient(ctx context.Context) *gorm.DB {
	return DB.WithContext(ctx)
}
