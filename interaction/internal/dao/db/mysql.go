package db

import (
	"fmt"
	"github.com/sunflower10086/TikTok/interaction/internal/config"
	"github.com/sunflower10086/TikTok/interaction/internal/dao/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	db *gorm.DB
)

func GetDB() *gorm.DB {
	return db
}

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.C().MySQL.User,
		config.C().MySQL.Password,
		config.C().MySQL.Host,
		config.C().MySQL.Port,
		config.C().MySQL.Dbname,
	)

	// 日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	sqlDB, _ := _db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.C().MySQL.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.C().MySQL.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return err
	}

	db = _db

	err = autoMigrateDB(_db)
	if err != nil {
		return err
	}

	return nil
}

//func Init(c *config.Config) error {
//	if c == nil {
//		fmt.Println("配置文件对象为空")
//	} else {
//		fmt.Printf("config: %+v\n", c)
//	}
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		c.MySQL.User,
//		c.MySQL.Password,
//		c.MySQL.Host,
//		c.MySQL.Port,
//		c.MySQL.Dbname,
//	)
//
//	// 日志配置
//	newLogger := logger.New(
//		log.New(os.Stdout, "\n", log.LstdFlags),
//		logger.Config{
//			SlowThreshold:             time.Second,
//			Colorful:                  true,
//			IgnoreRecordNotFoundError: false,
//			ParameterizedQueries:      false,
//			LogLevel:                  logger.Info,
//		},
//	)
//
//	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		Logger: newLogger,
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//	})
//	if err != nil {
//		return err
//	}
//
//	sqlDB, _ := _db.DB()
//
//	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
//	sqlDB.SetMaxIdleConns(c.MySQL.MaxIdleConns)
//
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	sqlDB.SetMaxOpenConns(c.MySQL.MaxOpenConns)
//
//	// SetConnMaxLifetime 设置了连接可复用的最大时间。
//	sqlDB.SetConnMaxLifetime(time.Hour)
//	if err != nil {
//		return err
//	}
//
//	db = _db
//
//	err = autoMigrateDB(_db)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// 自动迁移数据库，如果没有表则自动创建
func autoMigrateDB(db *gorm.DB) error {
	// 创建User表
	err := db.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{})

	if err != nil {
		return err
	}

	return nil
}
