package controllers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//  "gin-go-test/routes"
	"gin-go-test/utils"
)

func TestImportStudentsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/api/import_students", ImportStudentsHandler)

	// Assume this file exists and is a valid test CSV
	csvPath := filepath.Join(utils.ProjectRoot(), "static", "exports", "score_0519103002_30cd.csv")

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	file, err := os.Open(csvPath)
	if err != nil {
		t.Fatalf("❌ Failed to open test CSV file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(csvPath))
	if err != nil {
		t.Fatalf("❌ Failed to create form file: %v", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("❌ Failed to copy file contents: %v", err)
	}

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/api/import_students", &buffer)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	t.Logf("✅ Response body: %s", resp.Body.String())
}
