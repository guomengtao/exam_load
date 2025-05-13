package utils

import (
	"fmt"
	"log"
	"os"
	"gin-go-test/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GormDB *gorm.DB

func InitGorm() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("MYSQL_USER", "root"),
		config.GetEnv("MYSQL_PASSWORD", ""),
		config.GetEnv("MYSQL_HOST", "127.0.0.1"),
		config.GetEnv("MYSQL_PORT", "3306"),
		config.GetEnv("MYSQL_DB", "test"),
	)

	tablePrefix := os.Getenv("TABLE_PREFIX")
	if tablePrefix == "" {
		tablePrefix = "tm_"
	}

	var err error
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 这里加前缀
			SingularTable: true,        // 表名不加复数
		},
	})
	if err != nil {
		log.Fatal("❌ GORM 初始化失败：", err)
	}
	fmt.Println("✅ GORM connected")
}