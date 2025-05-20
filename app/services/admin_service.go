package services

import (
	"fmt"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"errors"
	
 )

// GetAllAdmins retrieves all admin users from the database.
// The controller should wrap this result in a RESTful format like:
// { "code": 200, "message": "Success", "data": { "items": [...], "total": N } }
func GetAllAdmins() ([]models.Admin, error) {
	var admins []models.Admin

	table := utils.PrefixTable("admin")
	query := "SELECT * FROM " + table
	err := utils.DBX.Select(&admins, query)
	return admins, err
}

// UpdateOwnPassword updates the password for the admin identified by token-resolved adminID.
func UpdateOwnPassword(newPassword string, tokenString string) error {
	claims, err := utils.GetJWTInfo(tokenString)
	if err != nil {
		return errors.New("token 解析失败")
	}
	fmt.Println("JWT claims:", claims)

	adminIDFloat, ok := claims["admin_id"].(float64)
	if !ok {
		return errors.New("token 中未找到有效的管理员 ID")
	}
	adminID := int(adminIDFloat)

	hashedPassword := utils.HashPassword(newPassword)
	query := "UPDATE " + utils.PrefixTable("admin") + " SET password = ? WHERE id = ?"
	result, err := utils.DBX.Exec(query, hashedPassword, adminID)
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

func GetAdminByUsername(username string) (*models.Admin, error) {
	fmt.Println("🔍 准备查询用户名:", username)

	var admin models.Admin
	query := "SELECT id, username, password, role_id FROM " + utils.PrefixTable("admin") + " WHERE username = ? LIMIT 1"
	fmt.Println("📄 SQL 查询语句:", query)

	err := utils.DBX.Get(&admin, query, username)
	if err != nil {
		fmt.Println("❌ 查询失败:", err)
		return nil, err
	}

	fmt.Println("✅ 查询成功:", admin)
	return &admin, nil
}