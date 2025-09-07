// models/comment.go
package models

import (
	"time"
)

// Comment 评论模型
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content" binding:"required"`
	AuthorID  uint      `gorm:"not null" json:"author_id"`
	Author    User      `gorm:"foreignKey:AuthorID" json:"author"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CommentCreateRequest 创建评论请求
type CommentCreateRequest struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// CommentResponse 评论响应
type CommentResponse struct {
	ID        uint         `json:"id"`
	Content   string       `json:"content"`
	Author    UserResponse `json:"author"`
	PostID    uint         `json:"post_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
