package models

import (
	"time"
)

// User 数据模型  Key
type User struct {
	Id          int       `gorm:"column:id" json:"id" validate:"max=255"`
	Username    string    `gorm:"column:username" json:"username" validate:"required,max=255"`
	AdminId     string    `gorm:"column:admin_id" json:"admin_id" validate:"max=255"`
	Province    string    `gorm:"column:province" json:"province" validate:"max=255"`
	City        string    `gorm:"column:city" json:"city" validate:"max=255"`
	Area        string    `gorm:"column:area" json:"area" validate:"max=255"`
	SchoolName  string    `gorm:"column:school_name" json:"school_name" validate:"max=255"`
	GradeName   string    `gorm:"column:grade_name" json:"grade_name" validate:"max=255"`
	ClassName   string    `gorm:"column:class_name" json:"class_name" validate:"max=255"`
	UserId      string    `gorm:"column:user_id" json:"user_id" validate:"max=255"`
	Phone       string    `gorm:"column:phone" json:"phone" validate:"max=255"`
	UserType    bool      `gorm:"column:user_type" json:"user_type" validate:"max=255"`
	SchoolLevel bool      `gorm:"column:school_level" json:"school_level" validate:"max=255"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" validate:"max=255"`
}