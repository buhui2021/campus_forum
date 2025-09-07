// controllers/admin_controller.go
package controllers

import (
	"campus_forum/database"
	"campus_forum/models"
	"campus_forum/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPendingPosts 获取待审核帖子
func GetPendingPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var posts []models.Post
	if err := database.DB.Preload("Author").
		Where("status = ?", "pending").
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&posts).Error; err != nil {
		utils.InternalServerError(c, "获取待审核帖子失败")
		return
	}

	// 获取总数
	var total int64
	database.DB.Model(&models.Post{}).Where("status = ?", "pending").Count(&total)

	utils.Success(c, gin.H{
		"posts": posts,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// ApprovePost 审核通过帖子
func ApprovePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	if err := database.DB.Model(&post).Update("status", "approved").Error; err != nil {
		utils.InternalServerError(c, "审核帖子失败")
		return
	}

	utils.Success(c, gin.H{"message": "帖子审核通过"})
}

// RejectPost 审核拒绝帖子
func RejectPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	if err := database.DB.Model(&post).Update("status", "rejected").Error; err != nil {
		utils.InternalServerError(c, "审核帖子失败")
		return
	}

	utils.Success(c, gin.H{"message": "帖子审核拒绝"})
}

// PinPost 置顶帖子
func PinPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	// 切换置顶状态
	isPinned := !post.IsPinned
	if err := database.DB.Model(&post).Update("is_pinned", isPinned).Error; err != nil {
		utils.InternalServerError(c, "置顶帖子失败")
		return
	}

	action := "取消置顶"
	if isPinned {
		action = "置顶"
	}

	utils.Success(c, gin.H{"message": "帖子" + action + "成功", "is_pinned": isPinned})
}
