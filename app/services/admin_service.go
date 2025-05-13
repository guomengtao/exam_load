package services

import (
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"errors"
)

 

func GetAllAdmins() ([]models.Admin, error) {
	var admins []models.Admin
 
	table := utils.PrefixTable("admin")
	query := "SELECT * FROM " + table
	err := utils.DBX.Select(&admins, query)
	return admins, err
}

// 修改管理员密码
func UpdateAdminPassword(adminID int, newPassword string) error {
	query := "UPDATE " + utils.PrefixTable("admin") + " SET password = ? WHERE id = ?"
	result, err := utils.DBX.Exec(query, newPassword, adminID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("未找到该管理员")
	}
	return nil
}