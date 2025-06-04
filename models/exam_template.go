package models

import (
    "encoding/json"
    "time"
)

// ExamTemplate 试卷模板模型
type ExamTemplate struct {
    ID          int64           `json:"id" gorm:"primaryKey"`
    Title       string          `json:"title"`
    Description string          `json:"description"`
    CoverImage  string          `json:"cover_image"`
    TotalScore  int             `json:"total_score"`
    Questions   json.RawMessage `json:"questions"`
    CategoryID  int64           `json:"category_id"`
    PublishTime int             `json:"publish_time"`
    Status      int             `json:"status"`
    Creator     string          `json:"creator"`
    CreatedAt   time.Time       `json:"created_at"`
    UpdatedAt   time.Time       `json:"updated_at"`
}

// TableName 指定表名
func (ExamTemplate) TableName() string {
    return "tm_exam_template"
} 