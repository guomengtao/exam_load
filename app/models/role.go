// app/models/role.go
package models

type Role struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`  // 比如 admin、editor
	Desc string `json:"desc"`  // 描述：管理员、编辑者等
}