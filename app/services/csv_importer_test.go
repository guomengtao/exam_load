package services

import (
    "testing"
    "path/filepath"
    "gin-go-test/utils"
)

func TestImportStudentsFromCSV(t *testing.T) {
    filePath := filepath.Join(utils.ProjectRoot(), "static", "exports", "score_0519103002_30cd.csv") // Replace with actual test CSV path

    // Call the import function
    count, skipped, err := ImportStudentsFromCSV(filePath)
    if err != nil {
        t.Errorf("❌ CSV import failed: %v", err)
    } else {
        t.Logf("✅ Successfully imported %d student records, skipped %d", count, skipped)
    }

    // Optional: Clean up or check DB state if needed
}