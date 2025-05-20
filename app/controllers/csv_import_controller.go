package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "path/filepath"
    "gin-go-test/app/services"
    "gin-go-test/utils"
)

// POST /api/import_students
func ImportStudentsHandler(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "No file uploaded"})
        return
    }

    // Save uploaded file to temp dir
    savePath := filepath.Join(utils.ProjectRoot(), "static", "uploads", file.Filename)
    if err := c.SaveUploadedFile(file, savePath); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to save file"})
        return
    }

    // Import the file
    count, skipped, err := services.ImportStudentsFromCSV(savePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Import failed", "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "Import successful",
        "imported": count,
        "skipped":  skipped,
    })
}