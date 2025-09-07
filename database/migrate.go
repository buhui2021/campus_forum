// database/migrate.go
package database

import (
	"campus_forum/models"
	"log"
)

// Migrate 自动迁移数据库表
func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.Like{}, // 确保包含Like表
	)
	if err != nil {
		log.Fatalf("❌ 数据库迁移失败: %v", err)
	}
	log.Println("✅ 数据库迁移成功")
}
