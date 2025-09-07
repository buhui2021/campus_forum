// models/like.go
package models

import (
	"time"
)

// Like 点赞模型
type Like struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	PostID    uint      `gorm:"not null" json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
