// controllers/auth_controller.go
package controllers

import (
	"campus_forum/database"
	"campus_forum/models"
	"campus_forum/utils"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var input models.UserRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	result := database.DB.Where("username = ?", input.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		utils.BadRequest(c, "用户名已存在")
		return
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		utils.InternalServerError(c, "密码加密失败")
		return
	}

	// 创建用户
	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
		Email:    input.Email,
		Role:     "student",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.InternalServerError(c, "创建用户失败")
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		utils.InternalServerError(c, "生成令牌失败")
		return
	}

	// 返回响应
	utils.Success(c, gin.H{
		"token": token,
		"user": models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	var input models.UserLoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.BadRequest(c, "请求参数错误")
		return
	}

	// 查找用户
	var user models.User
	result := database.DB.Where("username = ?", input.Username).First(&user)
	if result.RowsAffected == 0 {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 验证密码
	if !utils.CheckPassword(input.Password, user.Password) {
		utils.Unauthorized(c, "用户名或密码错误")
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		utils.InternalServerError(c, "生成令牌失败")
		return
	}

	// 返回响应
	utils.Success(c, gin.H{
		"token": token,
		"user": models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	})
}

// GetProfile 获取用户信息
func GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Avatar:   user.Avatar,
	})
}
