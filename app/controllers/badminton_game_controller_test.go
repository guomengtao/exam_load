package controllers

import (
	"bytes"
	"encoding/json"
	"gin-go-test/app/biz"
	"gin-go-test/app/models"
	"gin-go-test/app/services"
	"gin-go-test/utils/generated/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func setupTestRouter() (*gin.Engine, *gorm.DB, *services.BadmintonGameService) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to sqlite in-memory db: " + err.Error())
	}

	// 自动建表
	_ = db.AutoMigrate(&models.BadmintonGame{})

	service := services.NewBadmintonGameService(db)
	bizLayer := biz.NewBadmintonGameBiz(service)
	skeleton := controller.NewBadmintonGameSkeleton(bizLayer)

	r := gin.Default()
	r.PUT("/api/badminton_game", skeleton.BatchUpdateHandler)
	return r, db, service
}

func TestBatchUpdateController(t *testing.T) {
	r, db, _ := setupTestRouter()

	// Insert initial test data
	initialData := []models.BadmintonGame{
		{Id: ptrInt(1), Player1: ptrString("Tom"), Player2: ptrString("Jack"), Score1: ptrInt(21), Score2: ptrInt(22)},
		{Id: ptrInt(2), Player1: ptrString("Alice"), Player2: ptrString("Bob"), Score1: ptrInt(18), Score2: ptrInt(20)},
	}
	result := db.Create(&initialData)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(len(initialData)), result.RowsAffected)

	// Verify initial data saved correctly
	var savedData []models.BadmintonGame
	err := db.Find(&savedData).Error
	assert.NoError(t, err)
	assert.Equal(t, 2, len(savedData))
	fmt.Println("Initial inserted records:", savedData)

	// Prepare batch update payload
	updateItems := []map[string]interface{}{
		{"id": 1, "score2": 23},
		{"id": 2, "score1": 19},
	}
	body, _ := json.Marshal(updateItems)

	// Create HTTP request to update records
	req, _ := http.NewRequest(http.MethodPut, "/api/badminton_game", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the update request
	r.ServeHTTP(w, req)

	// Check response status code
	assert.Equal(t, 200, w.Code)

	// Parse response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify response content
	assert.Equal(t, float64(0), response["code"])
	assert.Equal(t, "批量更新成功", response["data"].(map[string]interface{})["message"])

	// Query updated records to verify changes
	var updatedRecords []models.BadmintonGame
	err = db.Find(&updatedRecords).Error
	assert.NoError(t, err)
	fmt.Println("Updated records:")
	for _, r := range updatedRecords {
		fmt.Printf("Id=%d, Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v\n",
			*r.Id,
			derefString(r.Player1),
			derefString(r.Player2),
			derefInt(r.Score1),
			derefInt(r.Score2),
			func(s *string) string {
				if s == nil {
					return "11999"
				}
				return *s
			}(r.Location),
		)
	}

	// Verify updated values
	for _, record := range updatedRecords {
		switch *record.Id {
		case 1:
			assert.Equal(t, 21, *record.Score1)
			assert.Equal(t, 23, *record.Score2)
		case 2:
			assert.Equal(t, 19, *record.Score1)
			assert.Equal(t, 20, *record.Score2)
		default:
			t.Errorf("Unexpected record Id %v found", record.Id)
		}
	}
}

func TestBatchUpdateWithoutPlayerFields(t *testing.T) {
	r, db, _ := setupTestRouter()

	// 插入初始数据
	initialData := []models.BadmintonGame{
		{Id: ptrInt(1), Player1: ptrString("Tom"), Player2: ptrString("Jack"), Score1: ptrInt(20), Score2: ptrInt(22)},
	}
	result := db.Create(&initialData)
	assert.NoError(t, result.Error)
	assert.Equal(t, int64(len(initialData)), result.RowsAffected)

	// 准备更新数据，只更新分数，不传 Player1, Player2
	updateItems := []map[string]interface{}{
		{"id": 1, "score1": 30, "score2": 28},
	}
	body, _ := json.Marshal(updateItems)

	req, _ := http.NewRequest(http.MethodPut, "/api/badminton_game", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// 查询更新后数据
	var updatedRecords []models.BadmintonGame
	err := db.Find(&updatedRecords).Error
	assert.NoError(t, err)
	assert.Equal(t, 1, len(updatedRecords))

	rec := updatedRecords[0]
	assert.Equal(t, "Tom", derefString(rec.Player1))
	assert.Equal(t, "Jack", derefString(rec.Player2))
	assert.Equal(t, 30, derefInt(rec.Score1))
	assert.Equal(t, 28, derefInt(rec.Score2))

	fmt.Printf("Updated record: Id=%d, Player1=%s, Player2=%s, Score1=%d, Score2=%d\n",
		derefInt(rec.Id), derefString(rec.Player1), derefString(rec.Player2), derefInt(rec.Score1), derefInt(rec.Score2))
}

func ptrInt(i int) *int {
	return &i
}

func derefString(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func derefInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func ptrString(s string) *string {
	return &s
}