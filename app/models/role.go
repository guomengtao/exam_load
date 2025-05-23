package models



// Role 数据模型
type Role struct {
	Id   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Desc string `gorm:"column:desc" json:"desc"`
}