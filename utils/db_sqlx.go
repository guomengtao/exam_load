package utils

import (
	"fmt"
	"gin-go-test/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DBX 是 sqlx 的数据库连接实例
var DBX *sqlx.DB

// TablePrefix 用于拼接表名前缀
var TablePrefix string

// InitDBX 初始化 sqlx 数据库连接并加载表前缀
func InitDBX() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetEnv("MYSQL_USER", "root"),
		config.GetEnv("MYSQL_PASSWORD", ""),
		config.GetEnv("MYSQL_HOST", "127.0.0.1"),
		config.GetEnv("MYSQL_PORT", "3306"),
		config.GetEnv("MYSQL_DB", "test"),
	)

	var err error
	DBX, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Failed to connect to MySQL:", err)
	}

	if err := DBX.Ping(); err != nil {
		log.Fatal("❌ MySQL ping error:", err)
	}

	// 读取并设置表前缀
	TablePrefix = os.Getenv("TABLE_PREFIX")
	if TablePrefix == "" {
		TablePrefix = "ym_" // 默认值
	}

	fmt.Println("✅ sqlx MySQL connected")
}

// PrefixTable 返回加上前缀后的表名
func PrefixTable(table string) string {
	return TablePrefix + table
}
