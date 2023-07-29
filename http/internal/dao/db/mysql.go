package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/sunflower10086/restful-api-demo/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	var d *gorm.DB
	once.Do(func() {
		d = db
	})
	return d
}

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.C().MySQL.UserName,
		conf.C().MySQL.Password,
		conf.C().MySQL.Host,
		conf.C().MySQL.Port,
		conf.C().MySQL.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(conf.C().MySQL.MaxIdleConn)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(conf.C().MySQL.MaxOpenConn)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return err
	}

	db = db

	return nil
}
