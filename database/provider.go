// database/provider.go
package database

import "gorm.io/gorm"

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
