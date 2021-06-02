package xmysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)

func NewDB() (dbHandle *gorm.DB, err error) {
	dbHandle, err = gorm.Open("mysql",  viper.GetString("db.local.dsn"))
	if err != nil {
		return
	}
	dbHandle.DB().SetMaxOpenConns(10)
	dbHandle.DB().SetMaxIdleConns(10)
	dbHandle.DB().SetConnMaxLifetime(1 * time.Minute)
	err = dbHandle.DB().Ping()
	return
}
