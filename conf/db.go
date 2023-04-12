package conf

import (
	"github.com/spf13/viper"
	"github.com/yaochen2019/gogofly/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open("db.dsn"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxIdleTime(time.Hour)
	db.AutoMigrate(&model.User{})
	return db, nil
}
