// routes/comments.go
package routes

import (
	"campus_forum/controllers"
	"campus_forum/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterCommentRoutes 注册评论路由
func RegisterCommentRoutes(r *gin.RouterGroup) {
	comments := r.Group("/comments")
	{
		comments.GET("/post/:id", controllers.GetComments)

		// 需要认证的路由
		authComments := comments.Use(middleware.AuthMiddleware())
		{
			authComments.POST("", controllers.CreateComment)
			authComments.DELETE("/:id", controllers.DeleteComment)
			authComments.GET("/my", controllers.GetUserComments) // 获取用户的所有评论
		}
	}
}
