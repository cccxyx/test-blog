package database

import (
	"log"

	"github.com/dingdinglz/test-blog/config"
	"github.com/dingdinglz/test-blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init() error {
	var err error

	// 获取数据库路径
	dbPath := config.AppConfig.Database.Path

	// 打开数据库连接
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Printf("数据库连接成功: %s\n", dbPath)

	// 自动迁移数据表
	err = DB.AutoMigrate(&models.User{}, &models.Article{})
	if err != nil {
		return err
	}

	log.Println("数据表迁移成功")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
