// models/post.go
package models

import (
	"time"
)

// Post 帖子模型
type Post struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"not null" json:"title" binding:"required"`
	Content    string    `gorm:"type:text;not null" json:"content" binding:"required"`
	AuthorID   uint      `gorm:"not null" json:"author_id"`
	Author     User      `gorm:"foreignKey:AuthorID" json:"author"`
	Status     string    `gorm:"default:'pending'" json:"status"` // pending, approved, rejected
	LikesCount int       `gorm:"default:0" json:"likes_count"`
	ViewsCount int       `gorm:"default:0" json:"views_count"`
	IsPinned   bool      `gorm:"default:false" json:"is_pinned"`
	Category   string    `json:"category"`
	Tags       string    `json:"tags"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Comments   []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

// PostCreateRequest 创建帖子请求
type PostCreateRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
}

// PostUpdateRequest 更新帖子请求
type PostUpdateRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Tags     string `json:"tags"`
}

// PostResponse 帖子响应
type PostResponse struct {
	ID         uint              `json:"id"`
	Title      string            `json:"title"`
	Content    string            `json:"content"`
	Author     UserResponse      `json:"author"`
	Status     string            `json:"status"`
	LikesCount int               `json:"likes_count"`
	ViewsCount int               `json:"views_count"`
	IsPinned   bool              `json:"is_pinned"`
	Category   string            `json:"category"`
	Tags       string            `json:"tags"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
	Comments   []CommentResponse `json:"comments,omitempty"`
}
