// ⚠️ 本文件为服务层模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=s 等命令重新生成覆盖。

package services

import (
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// {{.ModelName}}Service 处理{{.ModelName}}相关的业务逻辑
type {{.ModelName}}Service struct {
	db *gorm.DB
}

// New{{.ModelName}}Service 创建新的{{.ModelName}}Service实例
func New{{.ModelName}}Service(db *gorm.DB) *{{.ModelName}}Service {
	return &{{.ModelName}}Service{db: db}
}

// GetCount 获取记录总数
func (s *{{.ModelName}}Service) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.{{.ModelName}}{}).Count(&count).Error
	return count, err
}

// List 获取记录列表
func (s *{{.ModelName}}Service) List(page, pageSize int) ([]models.{{.ModelName}}, error) {
	var items []models.{{.ModelName}}
	err := s.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&items).Error
	return items, err
}

// BatchCreate 批量创建记录
func (s *{{.ModelName}}Service) BatchCreate(items []models.{{.ModelName}}) ([]models.{{.ModelName}}, []error) {
	var errs []error
	var createdItems []models.{{.ModelName}}

	for _, item := range items {
		if err := s.db.Create(&item).Error; err != nil {
			errs = append(errs, err)
		} else {
			createdItems = append(createdItems, item)
		}
	}

	return createdItems, errs
}

// BatchUpdate 批量更新记录
func (s *{{.ModelName}}Service) BatchUpdate(items []models.{{.ModelName}}) ([]models.{{.ModelName}}, []error) {
	var updatedItems []models.{{.ModelName}}
	var errs []error

	for _, item := range items {
		updateMap := make(map[string]interface{})
		{{- range .Fields }}
		{{- if ne .Name "Id" }}
		if item.{{ .GoName }} != nil {
			updateMap["{{ .Name }}"] = *item.{{ .GoName }}
		}
		{{- end }}
		{{- end }}
		if len(updateMap) > 0 {
			if err := s.db.Model(&models.{{.ModelName}}{}).Where("id = ?", *item.Id).Updates(updateMap).Error; err != nil {
				errs = append(errs, err)
				continue
			}
			var updatedItem models.{{.ModelName}}
			if err := s.db.First(&updatedItem, *item.Id).Error; err != nil {
				errs = append(errs, err)
				continue
			}
			updatedItems = append(updatedItems, updatedItem)
		}
	}
	return updatedItems, errs
}

// BatchDelete 批量删除记录
func (s *{{.ModelName}}Service) BatchDelete(ids []uint) []error {
	var errs []error

	for _, id := range ids {
		if err := s.db.Delete(&models.{{.ModelName}}{}, id).Error; err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func (s *{{.ModelName}}Service) GetDB() *gorm.DB {
	return s.db
}

func (s *{{.ModelName}}Service) GetDetail(id string) (*models.{{.ModelName}}, error) {
	var item models.{{.ModelName}}
	if err := s.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *{{.ModelName}}Service) ListWithOrder(page, pageSize int, sort, order string) ([]*models.{{.ModelName}}, []error) {
	var items []*models.{{.ModelName}}
	db := s.db.Model(&models.{{.ModelName}}{})
	if sort == "" {
		sort = "id"
	}
	if order == "" {
		order = "desc"
	}
	db = db.Order(sort + " " + order)
	err := db.Offset((page-1)*pageSize).Limit(pageSize).Find(&items).Error
	if err != nil {
		return nil, []error{err}
	}
	return items, nil
}