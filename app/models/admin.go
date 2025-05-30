package models

import (
	"time"
)

// Admin 数据模型  Key
type Admin struct {
	Id        int       `gorm:"column:id" json:"id" validate:"max=255"`
	Username  string    `gorm:"column:username" json:"username" validate:"required,max=255"`
	Password  string    `gorm:"column:password" json:"password" validate:"required,max=255"`
	RoleId    int       `gorm:"column:role_id" json:"role_id" validate:"required,max=255"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" validate:"max=255"`
}