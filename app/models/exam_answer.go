package models

import "time"

type ExamAnswer struct {
    AnswerUID   string    `gorm:"column:uuid;primaryKey"`     // 唯一作答记录标识（映射 Redis 的 answer_uid）
    ExamID      int       `gorm:"column:exam_id"`             // 试卷ID（int）
    ExamUUID    string    `gorm:"column:exam_uuid"`           // 试卷UUID
    UserID      string    `gorm:"column:user_id"`             // 学生编号
    Username    string    `gorm:"column:username"`            // 学生姓名
    Score       int       `gorm:"column:score"`               // 实际得分
    TotalScore  int       `gorm:"column:total_score"`         // 满分（int）
    Duration    int       `gorm:"column:duration"`            // 答题时长（秒）
    CreatedAt   time.Time `gorm:"column:created_at"`          // 创建时间
    Answers     string    `gorm:"column:answers;type:json"`   // JSON格式答题内容
}

// TableName 设置表名
func (ExamAnswer) TableName() string {
    return "tm_exam_answers"
}