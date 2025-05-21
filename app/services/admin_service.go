package services

import (
	"fmt"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"errors"
	"context"
	"time"
	
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

	if len(newPassword) < 6 {
		return errors.New("密码长度不能少于 6 位")
	}

	fmt.Println("🔐 开始加密密码...")
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("密码加密失败: " + err.Error())
	}
	fmt.Println("✅ 密码加密完成:", hashedPassword)

	adminIDFloat, ok := claims["admin_id"].(float64)
	if !ok {
		return errors.New("token 中未找到有效的管理员 ID")
	}
	adminID := int(adminIDFloat)

	var current models.Admin
	err = utils.DBX.Get(&current, "SELECT password FROM "+utils.PrefixTable("admin")+" WHERE id = ?", adminID)
	if err != nil {
		return errors.New("无法获取当前密码: " + err.Error())
	}
	fmt.Println("🧾 当前密码哈希:", current.Password)

	fmt.Println("🆕 将要写入的新密码哈希:", hashedPassword)

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
	// Refresh Redis cache with updated password
	var admin models.Admin
	query = "SELECT id, username, role_id FROM " + utils.PrefixTable("admin") + " WHERE id = ?"
	err = utils.DBX.Get(&admin, query, adminID)
	if err != nil {
		return errors.New("获取管理员信息失败")
	}

	cacheKey := "admin:" + admin.Username
	utils.RedisClient.Del(context.Background(), cacheKey)
	utils.RedisClient.HSet(context.Background(), cacheKey, map[string]interface{}{
		"id":       admin.ID,
		"username": admin.Username,
		"password": hashedPassword,
		"role_id":  admin.RoleID,
	})
	utils.RedisClient.Expire(context.Background(), cacheKey, 24*time.Hour)

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