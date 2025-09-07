// routes/admin.go
package routes

import (
	"campus_forum/controllers"
	"campus_forum/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAdminRoutes 注册管理员路由
func RegisterAdminRoutes(r *gin.RouterGroup) {
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// 帖子管理
		admin.GET("/posts/pending", controllers.GetPendingPosts)
		admin.PUT("/posts/:id/approve", controllers.ApprovePost)
		admin.PUT("/posts/:id/reject", controllers.RejectPost)
		admin.PUT("/posts/:id/pin", controllers.PinPost)
	}
}
