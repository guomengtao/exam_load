package models

import (
	"time"
)

// FileInfo 数据模型  Key
type FileInfo struct {
	Id         int       `gorm:"column:id" json:"id" validate:"max=255"`
	FileName   string    `gorm:"column:file_name" json:"file_name" validate:"required,max=255"`
	FilePath   string    `gorm:"column:file_path" json:"file_path" validate:"max=255"`
	FileSize   int64     `gorm:"column:file_size" json:"file_size" validate:"max=255"`
	UploadedAt time.Time `gorm:"column:uploaded_at" json:"uploaded_at" validate:"max=255"`
}
