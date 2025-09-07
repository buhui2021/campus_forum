// database/seeds.go
package database

import (
	"campus_forum/config"
	"campus_forum/models"
	"campus_forum/utils"
	"log"
)

// Seed 填充初始数据
func Seed() {
	// 创建默认管理员账户
	adminPassword, _ := utils.HashPassword(config.AppConfig.AdminPassword)
	admin := models.User{
		Username: config.AppConfig.AdminUsername,
		Password: adminPassword,
		Email:    "admin@forum.edu",
		Role:     "admin",
	}

	result := DB.FirstOrCreate(&admin, models.User{Username: config.AppConfig.AdminUsername})
	if result.Error != nil {
		log.Printf("❌ 创建管理员账户失败: %v", result.Error)
	} else if result.RowsAffected > 0 {
		log.Println("✅ 默认管理员账户已创建")
	}

	// 创建一些示例学生账户
	students := []models.User{
		{
			Username: "student1",
			Password: adminPassword, // 实际中应该使用不同的密码
			Email:    "student1@forum.edu",
			Role:     "student",
		},
		{
			Username: "student2",
			Password: adminPassword,
			Email:    "student2@forum.edu",
			Role:     "student",
		},
	}

	for _, student := range students {
		result := DB.FirstOrCreate(&student, models.User{Username: student.Username})
		if result.Error != nil {
			log.Printf("❌ 创建学生账户失败: %v", result.Error)
		}
	}

	log.Println("✅ 初始数据填充完成")
}
