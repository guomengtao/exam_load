// ⚠️ 本文件为服务骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=s 等命令重新生成覆盖。

package service

import (
	"fmt"
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// NOTE: This template requires the 'camelCase' function to be registered when parsing the template.

type BadmintonGameServiceSkeleton struct {
	db *gorm.DB
}

func NewBadmintonGameServiceSkeleton(db *gorm.DB) *BadmintonGameServiceSkeleton {
	return &BadmintonGameServiceSkeleton{
		db: db,
	}
}

// GetCount 返回数据库中该模型的总记录数
func (s *BadmintonGameServiceSkeleton) GetCount() (int64, []ErrorResponse) {
	var count int64
	if err := s.db.Model(&models.BadmintonGame{}).Count(&count).Error; err != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "failed to get count", err.Error())}
	}
	return count, nil
}

// List 根据分页参数返回数据列表和总数
func (s *BadmintonGameServiceSkeleton) List(page int, pageSize int) ([]*models.BadmintonGame, int64, []ErrorResponse) {
	var items []*models.BadmintonGame
	var total int64

	offset := (page - 1) * pageSize

	err := s.db.Model(&models.BadmintonGame{}).
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
func (s *BadmintonGameServiceSkeleton) BatchCreate(items []*models.BadmintonGame) (int, []ErrorResponse) {
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
		fmt.Printf("Creating BadmintonGame: id = %v\n", item.Id)

		// player1 字段校验
		fmt.Printf("Creating BadmintonGame: player1 = %v\n", item.Player1)

		// player2 字段校验
		fmt.Printf("Creating BadmintonGame: player2 = %v\n", item.Player2)

		// score1 字段校验
		fmt.Printf("Creating BadmintonGame: score1 = %v\n", item.Score1)

		// score2 字段校验
		fmt.Printf("Creating BadmintonGame: score2 = %v\n", item.Score2)

		// location 字段校验
		fmt.Printf("Creating BadmintonGame: location = %v\n", item.Location)

		// match_time 字段校验
		fmt.Printf("Creating BadmintonGame: match_time = %v\n", item.MatchTime)

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
func (s *BadmintonGameServiceSkeleton) BatchUpdate(items []*models.BadmintonGame) (int, []ErrorResponse) {
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

		// player1 字段校验

		// player2 字段校验

		// score1 字段校验

		// score2 字段校验

		// location 字段校验

		// match_time 字段校验

		// 检查记录是否存在
		var existing models.BadmintonGame
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

		if item.Player1 != nil {
			updateMap["player1"] = *item.Player1
		}

		if item.Player2 != nil {
			updateMap["player2"] = *item.Player2
		}

		if item.Score1 != nil {
			updateMap["score1"] = *item.Score1
		}

		if item.Score2 != nil {
			updateMap["score2"] = *item.Score2
		}

		if item.Location != nil {
			updateMap["location"] = *item.Location
		}

		if item.MatchTime != nil {
			updateMap["match_time"] = *item.MatchTime
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
func (s *BadmintonGameServiceSkeleton) BatchDelete(ids []int64) (int, []ErrorResponse) {
	if len(ids) == 0 {
		return 0, []ErrorResponse{NewErrorResponse(400, "empty ids", "")}
	}

	// 开启事务
	tx := s.db.Begin()
	if tx.Error != nil {
		return 0, []ErrorResponse{NewErrorResponse(500, "begin transaction failed", tx.Error.Error())}
	}

	// 删除主表数据
	if err := tx.Delete(&models.BadmintonGame{}, ids).Error; err != nil {
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
