package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	. "timenoteWeb/log"
)

var logging = Logger.WithField("包", "database")

// DB 全局数据库对象
var DB *gorm.DB

// 初始化数据库
func init() {
	var log = logging.WithField("源", "init")
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.WithError(err).Fatal("打开数据库失败")
	}
	log.Info("打开数据库成功")
	err = DB.AutoMigrate(&Location{})
	if err != nil {
		log.WithError(err).Fatal("自动迁移数据库失败")
	}
}
