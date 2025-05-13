package services

import (
	"gin-go-test/app/models"
	"gin-go-test/utils"
)

type PaginatedRoles struct {
	List      []models.Role `json:"data"`
	Total     int64         `json:"-"`
	Page      int           `json:"-"`
	PageSize  int           `json:"-"`
}

// 获取分页角色
func GetRolesPaginated(page, pageSize int) (PaginatedRoles, error) {
	var roles []models.Role
	var total int64

	offset := (page - 1) * pageSize

	// 查询总数
	if err := utils.GormDB.Model(&models.Role{}).Count(&total).Error; err != nil {
		return PaginatedRoles{}, err
	}

	// 查询分页数据
	if err := utils.GormDB.
		Limit(pageSize).
		Offset(offset).
		Find(&roles).Error; err != nil {
		return PaginatedRoles{}, err
	}

	return PaginatedRoles{
		List:     roles,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}