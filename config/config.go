package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	projectRoot := findProjectRoot()
	envPath := filepath.Join(projectRoot, ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("⚠️ .env 加载失败: %v", err)
	} else {
		log.Printf("✅ 已加载 .env 配置: %s", envPath)
	}
}

// GetEnv 读取环境变量，如果没有则返回默认值
func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// 自动查找 go.mod 所在目录，作为项目根目录
func findProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "." // fallback
}
