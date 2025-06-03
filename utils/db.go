package utils

import (
	"database/sql"
	"fmt"
	"log"
	"gin-go-test/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	user := config.GetEnv("MYSQL_USER", "47_120_38_206")
	pass := config.GetEnv("MYSQL_PASSWORD", "2HPzxPm9dn")
	host := config.GetEnv("MYSQL_HOST", "47.120.38.206")
	port := config.GetEnv("MYSQL_PORT", "3306")
	dbname := config.GetEnv("MYSQL_DB", "47_120_38_206")

	log.Printf("🔍 正在连接数据库: %s@%s:%s/%s", user, host, port, dbname)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ MySQL 连接失败:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("❌ MySQL ping 失败:", err)
	}

	log.Println("✅ MySQL 连接成功")
}