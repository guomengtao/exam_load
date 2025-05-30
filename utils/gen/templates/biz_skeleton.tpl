// ⚠️ 本文件为骨架模板，禁止直接修改任何生成器生成的文件！
// 如需调整，请修改本模板，并通过 go run utils/gen/gen.go -table=表名 -cmd=b 等命令重新生成覆盖。

package biz

import (
	"gin-go-test/app/models"
	"gin-go-test/app/services"
)

// {{.TableName}}Biz 处理{{.TableName}}相关的业务逻辑
type {{.TableName}}Biz struct {
	service *services.{{.ModelName}}Service
}

// New{{.TableName}}Biz 创建新的{{.TableName}}Biz实例
func New{{.TableName}}Biz(service *services.{{.ModelName}}Service) *{{.TableName}}Biz {
	return &{{.TableName}}Biz{service: service}
}

// GetCount 获取记录总数
func (b *{{.TableName}}Biz) GetCount() (int64, error) {
	return b.service.GetCount()
}

// List 获取记录列表
func (b *{{.TableName}}Biz) List(page, pageSize int) ([]models.{{.ModelName}}, error) {
	return b.service.List(page, pageSize)
}

// BatchCreate 批量创建记录
func (b *{{.TableName}}Biz) BatchCreate(items []*models.{{.ModelName}}) ([]models.{{.ModelName}}, []error) {
	// 将 []*models.{{.ModelName}} 转换为 []models.{{.ModelName}}
	modelItems := make([]models.{{.ModelName}}, len(items))
	for i, item := range items {
		modelItems[i] = *item
	}
	return b.service.BatchCreate(modelItems)
}

// BatchUpdate 批量更新记录
func (b *{{.TableName}}Biz) BatchUpdate(items []*models.{{.ModelName}}) error {
	// 将 []*models.{{.ModelName}} 转换为 []models.{{.ModelName}}
	modelItems := make([]models.{{.ModelName}}, len(items))
	for i, item := range items {
		modelItems[i] = *item
	}
	return b.service.BatchUpdate(modelItems)
}

// BatchDelete 批量删除记录
func (b *{{.TableName}}Biz) BatchDelete(ids []int) []error {
	// 将 []int 转换为 []uint
	uintIds := make([]uint, len(ids))
	for i, id := range ids {
		uintIds[i] = uint(id)
	}
	return b.service.BatchDelete(uintIds)
}