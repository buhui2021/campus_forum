// repository/post_repository.go
package repository

import (
	"campus_forum/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// PostRepository 帖子数据访问层
type PostRepository struct {
	BaseRepository
}

// Create 创建帖子
func (r *PostRepository) Create(post *models.Post) error {
	return r.DB().Create(post).Error
}

// FindByID 根据ID查找帖子
func (r *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.DB().
		Preload("Author").
		Preload("Comments").
		Preload("Comments.Author").
		First(&post, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("post not found")
		}
		return nil, fmt.Errorf("database error: %v", err)
	}

	return &post, nil
}

// FindAll 查找所有帖子
func (r *PostRepository) FindAll(offset, limit int, status, category string) ([]models.Post, error) {
	var posts []models.Post
	query := r.DB().Preload("Author").Where("status = ?", status)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error
	return posts, err
}

// FindByUserID 根据用户ID查找帖子
func (r *PostRepository) FindByUserID(userID uint) ([]models.Post, error) {
	var posts []models.Post
	err := r.DB().
		Preload("Author").
		Where("author_id = ?", userID).
		Order("created_at DESC").
		Find(&posts).Error
	return posts, err
}

// FindPending 查找待审核的帖子
func (r *PostRepository) FindPending(offset, limit int) ([]models.Post, error) {
	var posts []models.Post
	err := r.DB().
		Preload("Author").
		Where("status = ?", "pending").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error
	return posts, err
}

// Update 更新帖子
func (r *PostRepository) Update(post *models.Post) error {
	return r.DB().Save(post).Error
}

// Delete 删除帖子
func (r *PostRepository) Delete(post *models.Post) error {
	return r.DB().Delete(post).Error
}

// Count 统计帖子数量
func (r *PostRepository) Count(status, category string) (int64, error) {
	var count int64
	query := r.DB().Model(&models.Post{}).Where("status = ?", status)

	if category != "" {
		query = query.Where("category = ?", category)
	}

	err := query.Count(&count).Error
	return count, err
}

// CountPending 统计待审核帖子数量
func (r *PostRepository) CountPending() (int64, error) {
	var count int64
	err := r.DB().Model(&models.Post{}).Where("status = ?", "pending").Count(&count).Error
	return count, err
}
