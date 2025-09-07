// routes/auth.go
package routes

import (
	"campus_forum/controllers"
	"campus_forum/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes 注册认证路由
func RegisterAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)
	}
}
