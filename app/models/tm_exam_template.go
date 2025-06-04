package models

import (
	"time"
)

// TmExamTemplate 表示tm_exam_template表的结构
type TmExamTemplate struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	// 其他字段根据实际表结构添加
} 