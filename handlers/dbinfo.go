package handlers

import (
    "database/sql"
    "log"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

// DBInfo 用来存储每个字段的信息
type DBInfo struct {
    TableName    string `json:"table_name"`
    ColumnName   string `json:"column_name"`
    ColumnComment string `json:"column_comment"`
}

func GetDBTablesInfo(c *gin.Context) {
    // 连接数据库
    db, err := sql.Open("mysql", "user:password@tcp(47.120.38.206:3306)/47_120_38_206")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 查询表名和字段信息
    rows, err := db.Query(`
        SELECT 
            TABLE_NAME, 
            COLUMN_NAME, 
            COLUMN_COMMENT 
        FROM 
            INFORMATION_SCHEMA.COLUMNS 
        WHERE 
            TABLE_SCHEMA = '47_120_38_206' 
        ORDER BY 
            TABLE_NAME, ORDINAL_POSITION`)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // 存储查询结果
    var result []DBInfo
    for rows.Next() {
        var info DBInfo
        if err := rows.Scan(&info.TableName, &info.ColumnName, &info.ColumnComment); err != nil {
            log.Fatal(err)
        }
        result = append(result, info)
    }

    // 检查查询是否发生错误
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }

    // 返回查询结果为 JSON 格式
    c.JSON(200, result)
}