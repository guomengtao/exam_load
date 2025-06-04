// ⚠️ 本文件为服务骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=s 等命令重新生成覆盖。

package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// NOTE: This template requires the 'camelCase' function to be registered when parsing the template.

type ExamTemplateServiceSkeleton struct {
	db *gorm.DB
}

func NewExamTemplateServiceSkeleton(db *gorm.DB) *ExamTemplateServiceSkeleton {
	return &ExamTemplateServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *ExamTemplateServiceSkeleton) GetCount() (int64, []ErrorResponse) {
	var count int64
	if err := s.db.Model(&models.ExamTemplate{}).Count(&count).Error; err != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "failed to get count", err.Error())}
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *ExamTemplateServiceSkeleton) List(page int, pageSize int) ([]*models.ExamTemplate, int64, []ErrorResponse) {
	var items []*models.ExamTemplate
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.ExamTemplate{}).
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
func (s *ExamTemplateServiceSkeleton) BatchCreate(items []*models.ExamTemplate) (int, []ErrorResponse) {
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
		fmt.Printf("Creating ExamTemplate: id = %v\n", item.Id)
		
		// title 字段校验
		fmt.Printf("Creating ExamTemplate: title = %v\n", item.Title)
		
		// description 字段校验
		fmt.Printf("Creating ExamTemplate: description = %v\n", item.Description)
		
		// cover_image 字段校验
		fmt.Printf("Creating ExamTemplate: cover_image = %v\n", item.CoverImage)
		
		// total_score 字段校验
		fmt.Printf("Creating ExamTemplate: total_score = %v\n", item.TotalScore)
		
		// questions 字段校验
		fmt.Printf("Creating ExamTemplate: questions = %v\n", item.Questions)
		
		// category_id 字段校验
		fmt.Printf("Creating ExamTemplate: category_id = %v\n", item.CategoryId)
		
		// publish_time 字段校验
		fmt.Printf("Creating ExamTemplate: publish_time = %v\n", item.PublishTime)
		
		// status 字段校验
		fmt.Printf("Creating ExamTemplate: status = %v\n", item.Status)
		
		// creator 字段校验
		fmt.Printf("Creating ExamTemplate: creator = %v\n", item.Creator)
		
		// created_at 字段校验
		fmt.Printf("Creating ExamTemplate: created_at = %v\n", item.CreatedAt)
		
		// updated_at 字段校验
		fmt.Printf("Creating ExamTemplate: updated_at = %v\n", item.UpdatedAt)
		
		// deleted_at 字段校验
		fmt.Printf("Creating ExamTemplate: deleted_at = %v\n", item.DeletedAt)
		
		// time_limit 字段校验
		fmt.Printf("Creating ExamTemplate: time_limit = %v\n", item.TimeLimit)
		

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
func (s *ExamTemplateServiceSkeleton) BatchUpdate(items []*models.ExamTemplate) (int, []ErrorResponse) {
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
		if item.Id == nil || *item.Id <= 0 {
			errors = append(errors, NewErrorResponse(400, fmt.Sprintf("item[%d]: id is required", i), ""))
			continue
		}
		
		// id 字段校验
		
		
		// title 字段校验
		
		
		// description 字段校验
		
		
		// cover_image 字段校验
		
		
		// total_score 字段校验
		
		
		// questions 字段校验
		
		
		// category_id 字段校验
		
		
		// publish_time 字段校验
		
		
		// status 字段校验
		
		
		// creator 字段校验
		
		
		// created_at 字段校验
		
		
		// updated_at 字段校验
		
		
		// deleted_at 字段校验
		
		
		// time_limit 字段校验
		
		

		// 检查记录是否存在
		var existing models.ExamTemplate
		if err := tx.First(&existing, item.Id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				errors = append(errors, NewErrorResponse(404, fmt.Sprintf("item[%d]: record with id %d not found", i, item.Id), ""))
			} else {
				errors = append(errors, NewErrorResponse(500, fmt.Sprintf("item[%d]: check record exists failed", i), err.Error()))
			}
			continue
		}

		// 构建只包含非 nil 字段的更新 map
		updateMap := make(map[string]interface{})
		
		if item.Id != nil {
			updateMap["id"] = *item.Id
		}
		
		if item.Title != nil {
			updateMap["title"] = *item.Title
		}
		
		if item.Description != nil {
			updateMap["description"] = *item.Description
		}
		
		if item.CoverImage != nil {
			updateMap["cover_image"] = *item.CoverImage
		}
		
		if item.TotalScore != nil {
			updateMap["total_score"] = *item.TotalScore
		}
		
		if item.Questions != nil {
			updateMap["questions"] = *item.Questions
		}
		
		if item.CategoryId != nil {
			updateMap["category_id"] = *item.CategoryId
		}
		
		if item.PublishTime != nil {
			updateMap["publish_time"] = *item.PublishTime
		}
		
		if item.Status != nil {
			updateMap["status"] = *item.Status
		}
		
		if item.Creator != nil {
			updateMap["creator"] = *item.Creator
		}
		
		if item.CreatedAt != nil {
			updateMap["created_at"] = *item.CreatedAt
		}
		
		if item.UpdatedAt != nil {
			updateMap["updated_at"] = *item.UpdatedAt
		}
		
		if item.DeletedAt != nil {
			updateMap["deleted_at"] = *item.DeletedAt
		}
		
		if item.TimeLimit != nil {
			updateMap["time_limit"] = *item.TimeLimit
		}
		

		if len(updateMap) == 0 {
			errors = append(errors, NewErrorResponse(400, fmt.Sprintf("item[%d]: no fields to update", i), ""))
			continue
		}

		if err := tx.Model(&existing).Updates(updateMap).Error; err != nil {
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
func (s *ExamTemplateServiceSkeleton) BatchDelete(ids []int64) (int, []ErrorResponse) {
	if len(ids) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty ids", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	// 删除主表数据
	if err := tx.Delete(&models.ExamTemplate{}, ids).Error; err != nil {
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