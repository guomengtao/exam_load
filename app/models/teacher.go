package models

import (
	"time"
)

// Teacher 数据模型  Key
type Teacher struct {
	Id        int       `gorm:"column:id" json:"id" validate:"max=255"`
	Uuname    string    `gorm:"column:uuname" json:"uuname" validate:"required,max=255"`
	Email     string    `gorm:"column:email" json:"email" validate:"required,max=255"`
	Age       int       `gorm:"column:age" json:"age" validate:"max=255"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" validate:"max=255"`
	School    string    `gorm:"column:school" json:"school" validate:"required,max=255"`
}
