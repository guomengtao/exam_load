package meta

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

// GetDBFromEnv returns a *sql.DB using env variables
func GetDBFromEnv() (*sql.DB, error) {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")
	if user == "" { user = "root" }
	if pass == "" { pass = "123456" }
	if host == "" { host = "127.0.0.1" }
	if port == "" { port = "3306" }
	if dbname == "" { dbname = "gin_go_test" }

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)
	return sql.Open("mysql", dsn)
} 