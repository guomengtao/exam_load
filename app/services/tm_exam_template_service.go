package services

import (
	"gin-go-test/app/models"
	"gorm.io/gorm"
)

// TmExamTemplateService 处理tm_exam_template相关的服务逻辑
type TmExamTemplateService struct {
	db *gorm.DB
}

// NewTmExamTemplateService 创建新的TmExamTemplateService实例
func NewTmExamTemplateService(db *gorm.DB) *TmExamTemplateService {
	return &TmExamTemplateService{db: db}
}

// GetCount 获取记录总数
func (s *TmExamTemplateService) GetCount() (int64, error) {
	var count int64
	err := s.db.Model(&models.TmExamTemplate{}).Count(&count).Error
	return count, err
}

// List 获取记录列表
func (s *TmExamTemplateService) List(page, pageSize int) ([]models.TmExamTemplate, error) {
	var items []models.TmExamTemplate
	err := s.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&items).Error
	return items, err
}

// BatchCreate 批量创建记录
func (s *TmExamTemplateService) BatchCreate(items []models.TmExamTemplate) ([]models.TmExamTemplate, error) {
	err := s.db.Create(&items).Error
	return items, err
}

// BatchUpdate 批量更新记录
func (s *TmExamTemplateService) BatchUpdate(items []models.TmExamTemplate) error {
	for _, item := range items {
		if err := s.db.Save(&item).Error; err != nil {
			return err
		}
	}
	return nil
}

// BatchDelete 批量删除记录
func (s *TmExamTemplateService) BatchDelete(ids []uint) error {
	return s.db.Where("id IN ?", ids).Delete(&models.TmExamTemplate{}).Error
}

// GetDetail 获取记录详情
func (s *TmExamTemplateService) GetDetail(id string) (*models.TmExamTemplate, error) {
	var item models.TmExamTemplate
	err := s.db.First(&item, id).Error
	return &item, err
}

// ListWithOrder 获取记录列表（带排序）
func (s *TmExamTemplateService) ListWithOrder(page, pageSize int, sort, order string) ([]*models.TmExamTemplate, error) {
	var items []*models.TmExamTemplate
	query := s.db.Offset((page - 1) * pageSize).Limit(pageSize)
	if sort != "" {
		query = query.Order(sort + " " + order)
	}
	err := query.Find(&items).Error
	return items, err
} 