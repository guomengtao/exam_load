package meta

import (
	"database/sql"
	"fmt"
	"strings"
	
)

// GoName returns the Go-style (CamelCase) field name based on the database column name.
func (f Field) GoName() string {
	return toCamelCase(f.Name)
}

// toCamelCase converts snake_case to CamelCase (upper camel case).
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

// Field represents a database table column with metadata needed for Go struct generation.
type Field struct {
	Name           string // Column name
	Type           string // Go type corresponding to the column's data type
	IsRequired     bool   // Whether the column is NOT NULL
	IsPrimaryKey   bool   // Whether the column is a primary key
	IsAutoIncrement bool  // Whether the column is auto-incremented
}

// GoType returns the Go type as string (used for templates)
func (f Field) GoType() string {
	return f.Type
}

// CheckDBConnection checks if the database connection is alive and returns status
func CheckDBConnection(db *sql.DB) (bool, error) {
	// Ping the database to check connection
	err := db.Ping()
	if err != nil {
		return false, fmt.Errorf("database connection failed: %w", err)
	}
	return true, nil
}

// FetchTableFields queries the information_schema.columns table to retrieve metadata
// about the columns of the specified MySQL table and returns a slice of Field structs.
//
// Parameters:
//   - db: an open *sql.DB connection to the MySQL database
//   - tableName: the name of the table to fetch column information for
//
// Returns:
//   - []Field: slice of Field structs representing the columns of the table
//   - error: any error encountered during the query or scanning
func FetchTableFields(db *sql.DB, tableName string) ([]Field, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	if tableName == "" {
		return nil, fmt.Errorf("tableName is empty")
	}

	// Check database connection status
	isConnected, err := CheckDBConnection(db)
	if err != nil {
		return nil, fmt.Errorf("database connection check failed: %w", err)
	}
	if !isConnected {
		return nil, fmt.Errorf("database is not connected")
	}

	// Get current database name
	var dbName string
	err = db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to get current database: %w", err)
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ?", dbName, tableName).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("查询表是否存在失败: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("表 %s 不存在", tableName)
	}

	// Query to get column metadata from information_schema.columns
	const query = `
SELECT COLUMN_NAME, DATA_TYPE, IS_NULLABLE, COLUMN_KEY, EXTRA
FROM information_schema.columns
WHERE table_schema = ? AND table_name = ?
ORDER BY ORDINAL_POSITION
`

	rows, err := db.Query(query, dbName, tableName)
	if err != nil {
		return nil, fmt.Errorf("querying information_schema.columns failed: %w", err)
	}
	defer rows.Close()

	var fields []Field

	for rows.Next() {
		var columnName, dataType, isNullable, columnKey, extra string

		if err := rows.Scan(&columnName, &dataType, &isNullable, &columnKey, &extra); err != nil {
			return nil, fmt.Errorf("scanning row failed: %w", err)
		}

		field := Field{
			Name:           columnName,
			Type:           mapMySQLTypeToGoType(dataType),
			IsRequired:     strings.EqualFold(isNullable, "NO"),
			IsPrimaryKey:   columnKey == "PRI",
			IsAutoIncrement: strings.Contains(extra, "auto_increment"),
		}

		fields = append(fields, field)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	if len(fields) == 0 {
		return nil, fmt.Errorf("表 %s 不存在或没有字段", tableName)
	}

	// Remove debug print statements; return fields as is

	return fields, nil
}

// mapMySQLTypeToGoType maps common MySQL data types to Go types.
// This is a simplified mapping and may need extension for other types.
func mapMySQLTypeToGoType(mysqlType string) string {
	switch strings.ToLower(mysqlType) {
	case "char", "varchar", "text", "tinytext", "mediumtext", "longtext", "enum", "set":
		return "string"
	case "int", "integer", "smallint", "mediumint", "year":
		// Explicitly handle "mediumint" as int
		return "int"
	case "bigint":
		return "int64"
	case "tinyint":
		// tinyint(1) is often used as boolean, but here we treat all tinyint as int8 for simplicity
		return "int8"
	case "bit":
		return "[]byte"
	case "bool", "boolean":
		return "bool"
	case "float", "double", "real":
		return "float64"
	case "decimal", "numeric":
		return "string" // Could be decimal.Decimal or string, using string for simplicity
	case "date", "datetime", "timestamp", "time":
		return "time.Time"
	case "blob", "tinyblob", "mediumblob", "longblob", "binary", "varbinary":
		return "[]byte"
	default:
		// Unknown types default to string
		return "string"
	}
}