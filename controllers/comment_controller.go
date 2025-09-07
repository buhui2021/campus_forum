// controllers/comment_controller.go
package controllers

import (
	"campus_forum/models"
	"campus_forum/repository"
	"campus_forum/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var commentRepo = repository.CommentRepository{}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var input models.CommentCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 检查帖子是否存在
	var postRepo = repository.PostRepository{}
	_, err := postRepo.FindByID(input.PostID)
	if err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	// 创建评论
	comment := models.Comment{
		Content:  input.Content,
		AuthorID: userID.(uint),
		PostID:   input.PostID,
	}

	if err := commentRepo.Create(&comment); err != nil {
		utils.InternalServerError(c, "创建评论失败")
		return
	}

	// 获取完整的评论信息（包含作者信息）
	newComment, err := commentRepo.FindByID(comment.ID)
	if err != nil {
		utils.InternalServerError(c, "获取评论信息失败")
		return
	}

	utils.Success(c, newComment)
}

// GetComments 获取帖子评论
func GetComments(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	comments, err := commentRepo.FindByPostID(uint(postID))
	if err != nil {
		utils.InternalServerError(c, "获取评论失败")
		return
	}

	utils.Success(c, comments)
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "评论ID格式错误")
		return
	}

	// 查找评论
	comment, err := commentRepo.FindByID(uint(commentID))
	if err != nil {
		utils.NotFound(c, "评论不存在")
		return
	}

	// 检查权限（只有作者或管理员可以删除）
	if comment.AuthorID != userID.(uint) && role != "admin" {
		utils.Forbidden(c, "没有权限删除此评论")
		return
	}

	if err := commentRepo.Delete(comment); err != nil {
		utils.InternalServerError(c, "删除评论失败")
		return
	}

	utils.Success(c, gin.H{"message": "评论删除成功"})
}

// GetUserComments 获取用户的所有评论
func GetUserComments(c *gin.Context) {
	userID, _ := c.Get("user_id")

	comments, err := commentRepo.FindByUserID(userID.(uint))
	if err != nil {
		utils.InternalServerError(c, "获取用户评论失败")
		return
	}

	utils.Success(c, comments)
}
