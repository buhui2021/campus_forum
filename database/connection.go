// database/connection.go
package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	var err error
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	DB, err = gorm.Open(sqlite.Open("forum.db"), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		log.Fatalf("❌ 数据库连接失败: %v", err)
	}

	log.Println("✅ 数据库连接成功")
}

// CloseDB 关闭数据库连接
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("❌ 获取数据库实例失败: %v", err)
		return
	}
	sqlDB.Close()
	log.Println("✅ 数据库连接已关闭")
}
