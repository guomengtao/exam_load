package services

import (
    "testing"
)

func TestExportExamAnswersToCSV(t *testing.T) {
    examUUID := "40a89290-0ff9-4c01-a803-57155a24985c"  // Replace with actual exam UUID if needed
    school := ""                         // Empty string means no school filter
    limit := 1000
    offset := 0

    filePath, err := ExportAnswersToCSV(examUUID, school, limit, offset)
    if err != nil {
        t.Fatalf("❌ 导出失败: %v", err)
    }

    t.Logf("✅ 导出成功，文件路径: %s", filePath)
}