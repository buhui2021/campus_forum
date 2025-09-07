// repository/comment_repository.go
package repository

import (
	"campus_forum/models"
)

// CommentRepository 评论数据访问层
type CommentRepository struct {
	BaseRepository
}

// Create 创建评论
func (r *CommentRepository) Create(comment *models.Comment) error {
	return r.DB().Create(comment).Error
}

// FindByID 根据ID查找评论
func (r *CommentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	err := r.DB().First(&comment, id).Error
	return &comment, err
}

// FindByPostID 根据帖子ID查找评论
func (r *CommentRepository) FindByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB().
		Preload("Author").
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&comments).Error
	return comments, err
}

// FindByUserID 根据用户ID查找评论
func (r *CommentRepository) FindByUserID(userID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := r.DB().
		Preload("Post").
		Where("author_id = ?", userID).
		Order("created_at DESC").
		Find(&comments).Error
	return comments, err
}

// Update 更新评论
func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.DB().Save(comment).Error
}

// Delete 删除评论
func (r *CommentRepository) Delete(comment *models.Comment) error {
	return r.DB().Delete(comment).Error
}

// CountByPostID 统计帖子的评论数
func (r *CommentRepository) CountByPostID(postID uint) (int64, error) {
	var count int64
	err := r.DB().Model(&models.Comment{}).Where("post_id = ?", postID).Count(&count).Error
	return count, err
}
