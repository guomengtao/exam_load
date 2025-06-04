package meta

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// GenTable 表示需要生成代码的表信息
type GenTable struct {
	TableName    string `json:"table_name"`     // 表名
	CreatedAt    string `json:"created_at"`     // 创建时间
	UpdatedAt    string `json:"updated_at"`     // 更新时间
	GenerateCount int    `json:"generate_count"` // 生成次数
	IsDeleted    bool   `json:"is_deleted"`     // 是否已删除
}

// GenTableList 管理需要生成代码的表列表
type GenTableList struct {
	Tables []GenTable `json:"tables"`
	filePath string
}

// NewGenTableList 创建新的表列表管理器
func NewGenTableList() (*GenTableList, error) {
	// 确保配置文件目录存在
	configDir := filepath.Join("utils", "gen", "config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	filePath := filepath.Join(configDir, "gen_tables.data")
	list := &GenTableList{
		filePath: filePath,
	}

	// 如果文件存在，读取现有数据
	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(data, list); err != nil {
			return nil, err
		}
	}

	return list, nil
}

// Save 保存表列表到文件
func (l *GenTableList) Save() error {
	data, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(l.filePath, data, 0644)
}

// AddTable 添加新表到列表
func (l *GenTableList) AddTable(tableName string) error {
	// 检查表是否已存在
	for i, t := range l.Tables {
		if t.TableName == tableName {
			// 如果表被标记为删除，恢复它
			if t.IsDeleted {
				l.Tables[i].IsDeleted = false
				return l.Save()
			}
			return nil // 表已存在且未删除
		}
	}

	// 添加新表
	l.Tables = append(l.Tables, GenTable{
		TableName:    tableName,
		CreatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
		GenerateCount: 0,
		IsDeleted:    false,
	})

	return l.Save()
}

// RemoveTable 从列表中移除表（软删除）
func (l *GenTableList) RemoveTable(tableName string) error {
	for i, t := range l.Tables {
		if t.TableName == tableName {
			l.Tables[i].IsDeleted = true
			l.Tables[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return l.Save()
		}
	}
	return nil
}

// GetActiveTables 获取所有未删除的表
func (l *GenTableList) GetActiveTables() []string {
	var tables []string
	for _, t := range l.Tables {
		if !t.IsDeleted {
			tables = append(tables, t.TableName)
		}
	}
	return tables
}

// IncrementGenerateCount 增加表的生成次数
func (l *GenTableList) IncrementGenerateCount(tableName string) error {
	for i, t := range l.Tables {
		if t.TableName == tableName {
			l.Tables[i].GenerateCount++
			l.Tables[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			return l.Save()
		}
	}
	return nil
} 