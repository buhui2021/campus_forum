// config/config.go
package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig 存储应用程序配置
type Config struct {
	Environment   string
	DatabaseURL   string
	JWTSecret     string
	JWTTTL        int
	ServerPort    string
	AdminUsername string
	AdminPassword string
}

var AppConfig Config

// LoadConfig 从环境变量加载配置
func LoadConfig() {
	// 加载.env文件（如果存在）
	godotenv.Load()

	AppConfig = Config{
		Environment:   getEnv("ENVIRONMENT", "development"),
		DatabaseURL:   getEnv("DATABASE_URL", "forum.db"),
		JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
		JWTTTL:        getEnvAsInt("JWT_TTL", 24),
		ServerPort:    getEnv("PORT", "8080"),
		AdminUsername: getEnv("ADMIN_USERNAME", "admin"),
		AdminPassword: getEnv("ADMIN_PASSWORD", "admin123"),
	}

	log.Println("配置加载成功")
}

// 辅助函数
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt 从环境变量获取整数值
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
