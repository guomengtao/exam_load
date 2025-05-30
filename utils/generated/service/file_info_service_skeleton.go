package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// NOTE: This template requires the 'camelCase' function to be registered when parsing the template.

type FileInfoServiceSkeleton struct {
	db *gorm.DB
}

func NewFileInfoServiceSkeleton(db *gorm.DB) *FileInfoServiceSkeleton {
	return &FileInfoServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *FileInfoServiceSkeleton) GetCount() (int64, []ErrorResponse) {
	var count int64
	if err := s.db.Model(&models.FileInfo{}).Count(&count).Error; err != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "failed to get count", err.Error())}
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *FileInfoServiceSkeleton) List(page int, pageSize int) ([]*models.FileInfo, int64, []ErrorResponse) {
	var items []*models.FileInfo
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.FileInfo{}).
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
func (s *FileInfoServiceSkeleton) BatchCreate(items []*models.FileInfo) (int, []ErrorResponse) {
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
		fmt.Printf("Creating FileInfo: id = %v\n", item.Id)
		
		// file_name 字段校验
		fmt.Printf("Creating FileInfo: file_name = %v\n", item.FileName)
		
		// file_path 字段校验
		fmt.Printf("Creating FileInfo: file_path = %v\n", item.FilePath)
		
		// file_size 字段校验
		fmt.Printf("Creating FileInfo: file_size = %v\n", item.FileSize)
		
		// uploaded_at 字段校验
		fmt.Printf("Creating FileInfo: uploaded_at = %v\n", item.UploadedAt)
		

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
func (s *FileInfoServiceSkeleton) BatchUpdate(items []*models.FileInfo) (int, []ErrorResponse) {
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
		fmt.Printf("Updating FileInfo: id = %v\n", item.Id)
		
		// file_name 字段校验
		fmt.Printf("Updating FileInfo: file_name = %v\n", item.FileName)
		
		// file_path 字段校验
		fmt.Printf("Updating FileInfo: file_path = %v\n", item.FilePath)
		
		// file_size 字段校验
		fmt.Printf("Updating FileInfo: file_size = %v\n", item.FileSize)
		
		// uploaded_at 字段校验
		fmt.Printf("Updating FileInfo: uploaded_at = %v\n", item.UploadedAt)
		

		// 检查记录是否存在
		var existing models.FileInfo
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
func (s *FileInfoServiceSkeleton) BatchDelete(ids []int64) (int, []ErrorResponse) {
	if len(ids) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty ids", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	// 删除主表数据
	if err := tx.Delete(&models.FileInfo{}, ids).Error; err != nil {
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