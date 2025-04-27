package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// 自动加载 .env 文件
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found")
	}
}

// 读取环境变量，如果没有则返回默认值
func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}