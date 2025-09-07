// main.go
package main

import (
	"campus_forum/config"
	"campus_forum/database" // 导入迁移包
	"campus_forum/middleware"
	"campus_forum/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	database.InitDB()
	defer database.CloseDB()

	// 自动迁移数据库表
	database.Migrate()

	// 设置Gin模式
	if config.AppConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin路由器
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// 注册路由
	routes.RegisterRoutes(router)

	// 如果是开发环境，填充种子数据
	if config.AppConfig.Environment == "development" {
		database.Seed()
	}

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("校园论坛系统启动成功，端口: %s", port)
	log.Printf("环境: %s", config.AppConfig.Environment)
	log.Printf("启动时间: %s", time.Now().Format("2006-01-02 15:04:05"))

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}

}
