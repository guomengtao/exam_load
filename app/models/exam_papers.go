package models



// ExamPapers 数据模型  Key
type ExamPapers struct {
	Id *int64 `gorm:"column:id" json:"id" validate:"max=255"`
	Uuid *string `gorm:"column:uuid" json:"uuid" validate:"max=255"`
	TemplateId *int64 `gorm:"column:template_id" json:"template_id" validate:"max=255"`
	Title *string `gorm:"column:title" json:"title" validate:"required,max=255"`
	Description *string `gorm:"column:description" json:"description" validate:"max=255"`
	CoverImage *string `gorm:"column:cover_image" json:"cover_image" validate:"max=255"`
	TotalScore *int `gorm:"column:total_score" json:"total_score" validate:"required,max=255"`
	Questions *string `gorm:"column:questions" json:"questions" validate:"required,max=255"`
	ViewCount *int `gorm:"column:view_count" json:"view_count" validate:"max=255"`
	CategoryId *int64 `gorm:"column:category_id" json:"category_id" validate:"max=255"`
	PublishTime *int `gorm:"column:publish_time" json:"publish_time" validate:"max=255"`
	Status *bool `gorm:"column:status" json:"status" validate:"max=255"`
	Creator *string `gorm:"column:creator" json:"creator" validate:"max=255"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at" validate:"max=255"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at" validate:"max=255"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at" validate:"max=255"`
	TimeLimit *int `gorm:"column:time_limit" json:"time_limit" validate:"max=255"`
}