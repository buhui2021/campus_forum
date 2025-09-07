// routes/routes.go
package routes

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// API路由组
	api := r.Group("/api")
	{
		RegisterAuthRoutes(api)
		RegisterPostRoutes(api)
		RegisterCommentRoutes(api) // 注册评论路由
		RegisterAdminRoutes(api)
	}

	// 健康检查端点
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "校园论坛系统运行正常",
		})
	})
}
