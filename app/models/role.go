package models



// Role 数据模型  Key
type Role struct {
	Id   int    `gorm:"column:id" json:"id" validate:"max=255"`
	Name string `gorm:"column:name" json:"name" validate:"required,max=255"`
	Desc string `gorm:"column:desc" json:"desc" validate:"max=255"`
}