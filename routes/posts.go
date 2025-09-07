// routes/posts.go
package routes

import (
	"campus_forum/controllers"
	"campus_forum/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterPostRoutes 注册帖子路由
func RegisterPostRoutes(r *gin.RouterGroup) {
	posts := r.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)

		// 需要认证的路由
		authPosts := posts.Use(middleware.AuthMiddleware())
		{
			authPosts.POST("", controllers.CreatePost)
			authPosts.PUT("/:id", controllers.UpdatePost)
			authPosts.DELETE("/:id", controllers.DeletePost)
			authPosts.POST("/:id/like", controllers.LikePost)
			authPosts.GET("/:id/like", controllers.GetPostLikeStatus) // 添加获取点赞状态路由
		}
	}
}
