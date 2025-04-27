package utils

import (
	"database/sql"
	"fmt"
	"log"
	"gin-go-test/config" // 保持不变

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("MYSQL_USER", "root"),
		config.GetEnv("MYSQL_PASSWORD", ""),
		config.GetEnv("MYSQL_HOST", "127.0.0.1"),
		config.GetEnv("MYSQL_PORT", "3306"),
		config.GetEnv("MYSQL_DB", "test"),
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("MySQL ping error:", err)
	}
	fmt.Println("✅ MySQL connected")
}