package models

import (
	"time"
)

// User 数据模型
type User struct {
	ID          string    `gorm:"column:id" json:"id"`
	Username    string    `gorm:"column:username" json:"username"`
	AdminID     string    `gorm:"column:admin_id" json:"admin_id"`
	Province    string    `gorm:"column:province" json:"province"`
	City        string    `gorm:"column:city" json:"city"`
	Area        string    `gorm:"column:area" json:"area"`
	SchoolName  string    `gorm:"column:school_name" json:"school_name"`
	GradeName   string    `gorm:"column:grade_name" json:"grade_name"`
	ClassName   string    `gorm:"column:class_name" json:"class_name"`
	UserID      string    `gorm:"column:user_id" json:"user_id"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	UserType    bool      `gorm:"column:user_type" json:"user_type"`
	SchoolLevel bool      `gorm:"column:school_level" json:"school_level"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}