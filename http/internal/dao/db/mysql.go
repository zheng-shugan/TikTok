package db

import (
	"fmt"
	"sync"
	"time"

	config "github.com/sunflower10086/TikTok/http/config"
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
		config.C().MySQL.User,
		config.C().MySQL.Password,
		config.C().MySQL.Host,
		config.C().MySQL.Port,
		config.C().MySQL.Dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.C().MySQL.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.C().MySQL.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return err
	}

	db = db

	return nil
}
