// controllers/post_controller.go
package controllers

import (
	"campus_forum/models"
	"campus_forum/repository"
	"campus_forum/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	postRepo = repository.PostRepository{}
	likeRepo = repository.LikeRepository{}
)

// CreatePost 创建帖子
func CreatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	var input models.PostCreateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 设置帖子状态（管理员发布的帖子自动审核通过）
	status := "pending"
	if role == "admin" {
		status = "approved"
	}

	// 创建帖子
	post := models.Post{
		Title:    input.Title,
		Content:  input.Content,
		AuthorID: userID.(uint),
		Status:   status,
		Category: input.Category,
		Tags:     input.Tags,
	}

	if err := postRepo.Create(&post); err != nil {
		utils.InternalServerError(c, "创建帖子失败")
		return
	}

	// 获取完整的帖子信息（包含作者信息）
	newPost, err := postRepo.FindByID(post.ID)
	if err != nil {
		utils.InternalServerError(c, "获取帖子信息失败")
		return
	}

	utils.Success(c, newPost)
}

// GetPosts 获取帖子列表
func GetPosts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.DefaultQuery("status", "approved")
	category := c.Query("category")

	offset := (page - 1) * limit

	posts, err := postRepo.FindAll(offset, limit, status, category)
	if err != nil {
		utils.InternalServerError(c, "获取帖子列表失败")
		return
	}

	// 获取总数
	total, err := postRepo.Count(status, category)
	if err != nil {
		utils.InternalServerError(c, "获取帖子总数失败")
		return
	}

	utils.Success(c, gin.H{
		"posts": posts,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	var postRepo = repository.PostRepository{}
	post, err := postRepo.FindByID(uint(id))
	if err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	// 增加浏览次数
	post.ViewsCount++
	if err := postRepo.Update(post); err != nil {
		utils.InternalServerError(c, "更新浏览量失败")
		return
	}

	utils.Success(c, post)
}

// LikePost 点赞帖子
func LikePost(c *gin.Context) {
	userID, _ := c.Get("user_id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	// 检查用户是否已经点赞过
	existingLike, err := likeRepo.FindByUserAndPost(userID.(uint), uint(id))

	// 如果已经点赞过，取消点赞
	if err == nil && existingLike != nil {
		if err := likeRepo.Delete(existingLike); err != nil {
			utils.InternalServerError(c, "取消点赞失败")
			return
		}

		// 更新帖子的点赞数
		post, err := postRepo.FindByID(uint(id))
		if err != nil {
			utils.InternalServerError(c, "获取帖子信息失败")
			return
		}

		post.LikesCount--
		if err := postRepo.Update(post); err != nil {
			utils.InternalServerError(c, "更新帖子点赞数失败")
			return
		}

		utils.Success(c, gin.H{
			"liked":       false,
			"likes_count": post.LikesCount,
		})
		return
	}

	// 如果没有点赞过，添加点赞
	like := models.Like{
		UserID: userID.(uint),
		PostID: uint(id),
	}

	if err := likeRepo.Create(&like); err != nil {
		utils.InternalServerError(c, "点赞失败")
		return
	}

	// 更新帖子的点赞数
	post, err := postRepo.FindByID(uint(id))
	if err != nil {
		utils.InternalServerError(c, "获取帖子信息失败")
		return
	}

	post.LikesCount++
	if err := postRepo.Update(post); err != nil {
		utils.InternalServerError(c, "更新帖子点赞数失败")
		return
	}

	utils.Success(c, gin.H{
		"liked":       true,
		"likes_count": post.LikesCount,
	})
}

// GetPostLikeStatus 获取用户对帖子的点赞状态
func GetPostLikeStatus(c *gin.Context) {
	userID, _ := c.Get("user_id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	// 检查用户是否已经点赞过
	existingLike, err := likeRepo.FindByUserAndPost(userID.(uint), uint(id))

	utils.Success(c, gin.H{
		"liked": err == nil && existingLike != nil,
	})
}

// 其他函数也需要类似地更新...

func UpdatePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	var input models.PostUpdateRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 查找帖子
	post, err := postRepo.FindByID(uint(id))
	if err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	// 检查权限（只有作者或管理员可以修改）
	if post.AuthorID != userID.(uint) && role != "admin" {
		utils.Forbidden(c, "没有权限修改此帖子")
		return
	}

	// 更新帖子
	if input.Title != "" {
		post.Title = input.Title
	}
	if input.Content != "" {
		post.Content = input.Content
	}
	if input.Category != "" {
		post.Category = input.Category
	}
	if input.Tags != "" {
		post.Tags = input.Tags
	}

	if err := postRepo.Update(post); err != nil {
		utils.InternalServerError(c, "更新帖子失败")
		return
	}

	utils.Success(c, post)
}

// DeletePost 删除帖子
func DeletePost(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.BadRequest(c, "帖子ID格式错误")
		return
	}

	// 查找帖子
	post, err := postRepo.FindByID(uint(id))
	if err != nil {
		utils.NotFound(c, "帖子不存在")
		return
	}

	// 检查权限（只有作者或管理员可以删除）
	if post.AuthorID != userID.(uint) && role != "admin" {
		utils.Forbidden(c, "没有权限删除此帖子")
		return
	}

	if err := postRepo.Delete(post); err != nil {
		utils.InternalServerError(c, "删除帖子失败")
		return
	}

	utils.Success(c, gin.H{"message": "帖子删除成功"})
}
