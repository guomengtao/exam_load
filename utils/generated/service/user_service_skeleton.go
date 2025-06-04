package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// NOTE: This template requires the 'camelCase' function to be registered when parsing the template.

type UserServiceSkeleton struct {
	db *gorm.DB
}

func NewUserServiceSkeleton(db *gorm.DB) *UserServiceSkeleton {
	return &UserServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *UserServiceSkeleton) GetCount() (int64, []ErrorResponse) {
	var count int64
	if err := s.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "failed to get count", err.Error())}
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *UserServiceSkeleton) List(page int, pageSize int) ([]*models.User, int64, []ErrorResponse) {
	var items []*models.User
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.User{}).
		Count(&total).
		Limit(pageSize).
		Offset(offset).
		Find(&items).Error
	if err != nil {
		return nil, 0, []ErrorResponse{NewErrorResponse(500, "failed to list items", err.Error())}
	}

	return items, total, nil
}

// BatchCreate 批量创建记录
func (s *UserServiceSkeleton) BatchCreate(items []*models.User) (int, []ErrorResponse) {
	if len(items) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty items", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	successCount := 0
	var errors []ErrorResponse

	for i, item := range items {

		// id 字段校验
		fmt.Printf("Creating User: id = %v\n", item.Id)

		// username 字段校验
		fmt.Printf("Creating User: username = %v\n", item.Username)

		// admin_id 字段校验
		fmt.Printf("Creating User: admin_id = %v\n", item.AdminId)

		// province 字段校验
		fmt.Printf("Creating User: province = %v\n", item.Province)

		// city 字段校验
		fmt.Printf("Creating User: city = %v\n", item.City)

		// area 字段校验
		fmt.Printf("Creating User: area = %v\n", item.Area)

		// school_name 字段校验
		fmt.Printf("Creating User: school_name = %v\n", item.SchoolName)

		// grade_name 字段校验
		fmt.Printf("Creating User: grade_name = %v\n", item.GradeName)

		// class_name 字段校验
		fmt.Printf("Creating User: class_name = %v\n", item.ClassName)

		// user_id 字段校验
		fmt.Printf("Creating User: user_id = %v\n", item.UserId)

		// phone 字段校验
		fmt.Printf("Creating User: phone = %v\n", item.Phone)

		// user_type 字段校验
		fmt.Printf("Creating User: user_type = %v\n", item.UserType)

		// school_level 字段校验
		fmt.Printf("Creating User: school_level = %v\n", item.SchoolLevel)

		// created_at 字段校验
		fmt.Printf("Creating User: created_at = %v\n", item.CreatedAt)

		// 创建记录
		if err := tx.Create(item).Error; err != nil {
			errors = append(errors, NewErrorResponse(500, fmt.Sprintf("item[%d]: create failed", i), err.Error()))
			continue
		}
		successCount++
	}

	// 如果有错误，回滚事务
	if len(errors) > 0 {
		if err := tx.Rollback().Error; err != nil {
			errors = append(errors, NewErrorResponse(500, "rollback failed", err.Error()))
		}
		return successCount, errors
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return successCount, append(errors, NewErrorResponse(500, "commit transaction failed", err.Error()))
	}

	return successCount, errors
}

// BatchUpdate 批量更新记录
func (s *UserServiceSkeleton) BatchUpdate(items []*models.User) (int, []ErrorResponse) {
	if len(items) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty items", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	successCount := 0
	var errors []ErrorResponse

	for i, item := range items {
		// 验证必填字段
		if item.Id <= 0 {
			errors = append(errors, NewErrorResponse(400, fmt.Sprintf("item[%d]: id is required", i), ""))
			continue
		}

		// id 字段校验
		fmt.Printf("Updating User: id = %v\n", item.Id)

		// username 字段校验
		fmt.Printf("Updating User: username = %v\n", item.Username)

		// admin_id 字段校验
		fmt.Printf("Updating User: admin_id = %v\n", item.AdminId)

		// province 字段校验
		fmt.Printf("Updating User: province = %v\n", item.Province)

		// city 字段校验
		fmt.Printf("Updating User: city = %v\n", item.City)

		// area 字段校验
		fmt.Printf("Updating User: area = %v\n", item.Area)

		// school_name 字段校验
		fmt.Printf("Updating User: school_name = %v\n", item.SchoolName)

		// grade_name 字段校验
		fmt.Printf("Updating User: grade_name = %v\n", item.GradeName)

		// class_name 字段校验
		fmt.Printf("Updating User: class_name = %v\n", item.ClassName)

		// user_id 字段校验
		fmt.Printf("Updating User: user_id = %v\n", item.UserId)

		// phone 字段校验
		fmt.Printf("Updating User: phone = %v\n", item.Phone)

		// user_type 字段校验
		fmt.Printf("Updating User: user_type = %v\n", item.UserType)

		// school_level 字段校验
		fmt.Printf("Updating User: school_level = %v\n", item.SchoolLevel)

		// created_at 字段校验
		fmt.Printf("Updating User: created_at = %v\n", item.CreatedAt)

		// 检查记录是否存在
		var existing models.User
		if err := tx.First(&existing, item.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors = append(errors, NewErrorResponse(404, fmt.Sprintf("item[%d]: record with id %d not found", i, item.Id), ""))
			} else {
				errors = append(errors, NewErrorResponse(500, fmt.Sprintf("item[%d]: check record exists failed", i), err.Error()))
			}
			continue
		}

		// 更新记录
		if err := tx.Model(&existing).Updates(item).Error; err != nil {
			errors = append(errors, NewErrorResponse(500, fmt.Sprintf("item[%d]: update failed", i), err.Error()))
			continue
		}
		successCount++
	}

	// 如果有错误，回滚事务
	if len(errors) > 0 {
		if err := tx.Rollback().Error; err != nil {
			errors = append(errors, NewErrorResponse(500, "rollback failed", err.Error()))
		}
		return successCount, errors
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return successCount, append(errors, NewErrorResponse(500, "commit transaction failed", err.Error()))
	}

	return successCount, errors
}

// BatchDelete 批量删除记录
func (s *UserServiceSkeleton) BatchDelete(ids []int64) (int, []ErrorResponse) {
	if len(ids) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty ids", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	// 删除主表数据
	if err := tx.Delete(&models.User{}, ids).Error; err != nil {
		if err := tx.Rollback().Error; err != nil {
			return 0, []ErrorResponse{NewErrorResponse(500, "rollback failed", err.Error())}
		}
		return 0, []ErrorResponse{NewErrorResponse(500, "delete records failed", err.Error())}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "commit transaction failed", err.Error())}
	}

	return len(ids), nil
}
