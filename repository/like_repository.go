// repository/like_repository.go
package repository

import (
	"campus_forum/models"
)

// LikeRepository 点赞数据访问层
type LikeRepository struct {
	BaseRepository
}

// Create 创建点赞记录
func (r *LikeRepository) Create(like *models.Like) error {
	return r.DB().Create(like).Error
}

// Delete 删除点赞记录
func (r *LikeRepository) Delete(like *models.Like) error {
	return r.DB().Delete(like).Error
}

// FindByUserAndPost 根据用户和帖子查找点赞记录
func (r *LikeRepository) FindByUserAndPost(userID, postID uint) (*models.Like, error) {
	var like models.Like
	err := r.DB().Where("user_id = ? AND post_id = ?", userID, postID).First(&like).Error
	return &like, err
}

// CountByPost 统计帖子的点赞数
func (r *LikeRepository) CountByPost(postID uint) (int64, error) {
	var count int64
	err := r.DB().Model(&models.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}

// FindByUserID 根据用户ID查找点赞记录
func (r *LikeRepository) FindByUserID(userID uint) ([]models.Like, error) {
	var likes []models.Like
	err := r.DB().
		Preload("Post").
		Preload("Post.Author").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&likes).Error
	return likes, err
}
