package services

import (
    "encoding/csv"
    "fmt"
    "os"
    "strings"
 
    "gin-go-test/utils"
)

func ImportStudentsFromCSV(filePath string) (int, int, error) {
    db := utils.DBX
    file, err := os.Open(filePath)
    if err != nil {
        return 0, 0, fmt.Errorf("无法打开文件: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.FieldsPerRecord = -1

    records, err := reader.ReadAll()
    if err != nil {
        return 0, 0, fmt.Errorf("读取CSV失败: %v", err)
    }

    if len(records) < 2 {
        return 0, 0, fmt.Errorf("CSV文件没有有效数据")
    }

    inserted := 0
    skipped := 0

    for i, row := range records {
        if i == 0 {
            continue // skip header
        }
        if len(row) < 2 {
            skipped++
            continue
        }

        userID := strings.TrimSpace(row[0])
        username := strings.TrimSpace(row[1])

        if userID == "" || username == "" {
            skipped++
            continue
        }

        var exists int
        err = db.Get(&exists, "SELECT COUNT(1) FROM tm_user WHERE user_id = ?", userID)
        if err != nil {
            skipped++
            continue
        }

        if exists > 0 {
            skipped++
            continue
        }

        _, err = db.Exec(`INSERT INTO tm_user (user_id, username) VALUES (?, ?)`, userID, username)
        if err != nil {
            skipped++
            continue
        }

        inserted++
    }

    return inserted, skipped, nil
}