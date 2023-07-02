package config

import (
	"fmt"
	"gogofly/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDb() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
  fmt.Println(viper.GetString("db.dsn"))
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	// 配置数据库参数
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	sqlDb.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	sqlDb.SetConnMaxLifetime(time.Hour)

  // 表结构schema迁移(同步)
  db.AutoMigrate(&model.User{})

	return db, nil
}
