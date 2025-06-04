package genlib

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/go-sql-driver/mysql"
)

type Field struct {
	Name            string
	Type            string
	Column          string
	JSON            string
	IsRequired      bool
	IsPrimaryKey    bool
	IsAutoIncrement bool
}

type ModelTemplateData struct {
	ModelName string
	Fields    []Field
	Imports   []string
}

var typeMap = map[string]string{
	"int":       "int",
	"mediumint": "int",
	"bigint":    "int64",
	"varchar":   "string",
	"text":      "string",
	"datetime":  "time.Time",
	"timestamp": "time.Time",
	"tinyint":   "bool",
}

func mapType(sqlType string) string {
	for key, goType := range typeMap {
		if strings.HasPrefix(sqlType, key) {
			return goType
		}
	}
	return "string"
}

func FetchColumns(db *sql.DB, table string) ([]Field, error) {
	rows, err := db.Query(fmt.Sprintf("SHOW COLUMNS FROM `%s`", table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fields []Field
	for rows.Next() {
		var field, typ, null, key, extra string
		var def sql.NullString
		if err := rows.Scan(&field, &typ, &null, &key, &def, &extra); err != nil {
			return nil, err
		}
		fields = append(fields, Field{
			Name:            toCamelCase(field),
			Type:            mapType(typ),
			Column:          field,
			JSON:            field,
			IsRequired:      null == "NO",
			IsPrimaryKey:    key == "PRI",
			IsAutoIncrement: strings.Contains(extra, "auto_increment"),
		})
	}
	return fields, nil
}

func GenerateModelFromTable(tableName string) error {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")

	tablePrefix := os.Getenv("TABLE_PREFIX")
	fullTableName := tableName
	if tablePrefix != "" {
		fullTableName = tablePrefix + tableName
	}

	fmt.Printf("📡 已经连接数据库: %s:%s/%s\n", host, port, dbname)

	fmt.Printf("🔍 获取到表结构: %s\n", fullTableName)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	fields, err := FetchColumns(db, fullTableName)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1146 {
			return fmt.Errorf("❌ 查询表结构失败: 表 %s 不存在", fullTableName)
		}
		return fmt.Errorf("查询表结构失败: %v", err)
	}

	modelName := toCamelCase(tableName)

	data := ModelTemplateData{
		ModelName: modelName,
		Fields:    fields,
		Imports:   []string{},
	}

	tmpl, err := template.ParseFiles("utils/gen/templates/model.tpl")
	if err != nil {
		return fmt.Errorf("加载模板失败: %v", err)
	}

	outputPath := fmt.Sprintf("app/models/%s.go", tableName)
	if err := os.MkdirAll("app/models", os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建模型文件失败: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("渲染模板失败: %v", err)
	}

	fmt.Println("✅ 模型生成成功:", outputPath)
	return nil
}

// FetchTableFields is a wrapper for FetchColumns to fetch table fields.
func FetchTableFields(db *sql.DB, tableName string) ([]Field, error) {
	return FetchColumns(db, tableName)
}
