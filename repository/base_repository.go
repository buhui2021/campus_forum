// repository/base_repository.go
package repository

import (
	"campus_forum/database"

	"gorm.io/gorm"
)

// BaseRepository 基础存储库，提供通用的数据库操作方法
type BaseRepository struct{}

// DB 获取数据库实例
func (r *BaseRepository) DB() *gorm.DB {
	return database.GetDB()
}

// WithTransaction 执行事务操作
func (r *BaseRepository) WithTransaction(txFunc func(tx *gorm.DB) error) error {
	tx := r.DB().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // 重新抛出panic
		}
	}()

	if err := txFunc(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Create 创建记录
func (r *BaseRepository) Create(model interface{}) error {
	return r.DB().Create(model).Error
}

// Update 更新记录
func (r *BaseRepository) Update(model interface{}) error {
	return r.DB().Save(model).Error
}

// Delete 删除记录
func (r *BaseRepository) Delete(model interface{}) error {
	return r.DB().Delete(model).Error
}

// FindByID 根据ID查找记录
func (r *BaseRepository) FindByID(model interface{}, id uint) error {
	return r.DB().First(model, id).Error
}
