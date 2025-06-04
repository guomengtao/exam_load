package services

import (
	"crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gin-go-test/utils"
)

var ErrNoData = errors.New("没有对应数据，无法导出文件")

type ExportAnswerRow struct {
	StudentID   string    `db:"user_id"`
	StudentName string    `db:"username"`
	School      string    `db:"school_name"`
	Grade       string    `db:"grade_name"`
	Class       string    `db:"class_name"`
	ExamID      int       `db:"exam_id"`
	Score       int       `db:"score"`
	TotalScore  int       `db:"total_score"`
	CreatedAt   time.Time `db:"created_at"`
}

// ExportAnswersToCSV 根据 examUUID 和 school 条件查询导出 CSV 文件
func ExportAnswersToCSV(examUUID, school string, limit, offset int) (string, error) {
	baseDir := filepath.Join(utils.ProjectRoot(), "static", "exports")

	// 确保目录存在
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return "", fmt.Errorf("创建导出目录失败: %v", err)
	}

	// 生成随机后缀
	randBytes := make([]byte, 2)
	_, err := rand.Read(randBytes)
	if err != nil {
		return "", fmt.Errorf("生成随机后缀失败: %v", err)
	}
	suffix := hex.EncodeToString(randBytes)

	// 生成带随机后缀的文件名（保持 .csv 后缀）
	filename := fmt.Sprintf("score_%s_%s.csv", utils.NowTimeString(), suffix)
	filePath := filepath.Join(baseDir, filename)

	db := utils.DBX // sqlx.DB

	// 基础查询语句
	query := `
        SELECT
            u.user_id, u.username, u.school_name, u.grade_name, u.class_name,
            a.exam_id, a.score, a.total_score, a.created_at
        FROM tm_exam_answers a
        LEFT JOIN tm_user u ON a.user_id = u.user_id
        WHERE a.exam_uuid = ?
    `

	// 参数切片
	args := []interface{}{examUUID}

	// 如果 school 不为空，添加条件
	if school != "" {
		query += " AND u.school_name = ? "
		args = append(args, school)
	}

	query += " ORDER BY a.id ASC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	var rows []ExportAnswerRow
	err = db.Select(&rows, query, args...)
	if err != nil {
		return "", fmt.Errorf("查询导出数据失败: %v", err)
	}

	if len(rows) == 0 {
		return "", ErrNoData
	}

	// 创建 CSV 文件
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建导出文件失败: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	err = writer.Write([]string{"学号", "姓名", "学校", "年级", "班级", "考试ID", "得分", "总分", "作答时间"})
	if err != nil {
		log.Printf("写入CSV表头失败: %v", err)
	}

	for _, row := range rows {
		createdTime := row.CreatedAt.Format("2006-01-02 15:04:05")
		err := writer.Write([]string{
			row.StudentID,
			row.StudentName,
			row.School,
			row.Grade,
			row.Class,
			strconv.Itoa(row.ExamID),
			strconv.Itoa(row.Score),
			strconv.Itoa(row.TotalScore),
			createdTime,
		})
		if err != nil {
			log.Printf("写入CSV记录失败: %v", err)
		}
	}

	fmt.Printf("✅ 导出成功，共 %d 条记录，文件路径: %s\n", len(rows), filePath)
	return filePath, nil
}
