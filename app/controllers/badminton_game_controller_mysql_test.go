package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gin-go-test/app/models"
	"gin-go-test/app/services"
	"gin-go-test/app/biz"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gin-go-test/utils/generated/controller"
)

// setupMySQLTestDB initializes and returns a *gorm.DB connected to MySQL for tests
func setupMySQLTestDB() *gorm.DB {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DB")

	if user == "" || password == "" || host == "" || port == "" || database == "" {
		log.Fatal("MySQL connection environment variables are not fully set")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", user, password, host, port, database)
	log.Printf("Connecting to MySQL test database at %s:%s (db: %s)", host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect MySQL test database: %v", err)
	}

	// Migrate the schema for BadmintonGame (optional, ensure table exists)
	err = db.AutoMigrate(&models.BadmintonGame{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	return db
}

func setupRouterWithMySQL(db *gorm.DB) *gin.Engine {
	r := gin.New()
	bizLayer := biz.NewBadmintonGameBiz(services.NewBadmintonGameService(db))
	skeleton := controller.NewBadmintonGameSkeleton(bizLayer)

	group := r.Group("/api/badminton_game")
	group.PUT("", skeleton.BatchUpdateHandler)
	group.POST("", skeleton.BatchCreateHandler)
	group.GET("/list", skeleton.ListHandler)
	group.GET("/count", skeleton.CountHandler)
	group.DELETE("", skeleton.BatchDeleteHandler)

	return r
}

func TestBatchUpdateControllerMySQL(t *testing.T) {
	log.Println("Starting MySQL integration test for BatchUpdateController")
	db := setupMySQLTestDB()

	// Clear existing test data
	log.Println("Clearing existing badminton_games table")
	db.Exec("DELETE FROM badminton_games")

	router := setupRouterWithMySQL(db)

	// Insert initial records
	initialItems := []models.BadmintonGame{
		{
			Player1:  ptrString("Tom"),
			Player2:  ptrString("Jack"),
			Score1:   ptrInt(21),
			Score2:   ptrInt(15),
			Location: ptrString("Old Gym"),
		},
		{
			Player1:  ptrString("Alice"),
			Player2:  ptrString("Bob"),
			Score1:   ptrInt(18),
			Score2:   ptrInt(21),
			Location: ptrString("Old Gym"),
		},
	}

	for _, item := range initialItems {
		if err := db.Create(&item).Error; err != nil {
			t.Fatalf("failed to create initial record: %v", err)
		}
	}

	// Fetch inserted records to get IDs and log before update values
	var inserted []models.BadmintonGame
	if err := db.Find(&inserted).Error; err != nil {
		t.Fatalf("failed to fetch inserted records: %v", err)
	}
	assert.Equal(t, 2, len(inserted))

	for i, rec := range inserted {
		t.Logf("Before update record %d: Id=%v, Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v",
			i+1,
			rec.Id,
			valueOrNilString(rec.Player1),
			valueOrNilString(rec.Player2),
			valueOrNilInt(rec.Score1),
			valueOrNilInt(rec.Score2),
			valueOrNilString(rec.Location),
		)
	}

	// Prepare batch update: update Score1 and Score2 only for inserted IDs
	updatePayload := []map[string]interface{}{
		{
			"id":     inserted[0].Id,
			"score1": 25,
			"score2": 23,
		},
		{
			"id":     inserted[1].Id,
			"score1": 20,
			"score2": 22,
		},
	}

	// Convert updatePayload to JSON
	jsonData, err := json.Marshal(updatePayload)
	if err != nil {
		t.Fatalf("failed to marshal update payload: %v", err)
	}

	// Perform HTTP PUT request for batch update
	req := httptest.NewRequest("PUT", "/api/badminton_game", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	log.Println("Sending batch update PUT request")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Verify updated records from DB
	var updated []models.BadmintonGame
	if err := db.Find(&updated).Error; err != nil {
		t.Fatalf("failed to fetch updated records: %v", err)
	}

	for i, rec := range updated {
		t.Logf("After  update record %d: Id=%v, Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v",
			i+1,
			rec.Id,
			valueOrNilString(rec.Player1),
			valueOrNilString(rec.Player2),
			valueOrNilInt(rec.Score1),
			valueOrNilInt(rec.Score2),
			valueOrNilString(rec.Location),
		)
	}

	assert.Equal(t, 25, *updated[0].Score1)
	assert.Equal(t, 23, *updated[0].Score2)
	assert.Equal(t, "Tom", *updated[0].Player1) // Player1 unchanged
	assert.Equal(t, 20, *updated[1].Score1)
	assert.Equal(t, 22, *updated[1].Score2)
	assert.Equal(t, "Alice", *updated[1].Player1) // Player1 unchanged

	// Cleanup
	log.Println("Cleaning up: deleting all badminton_games")
	db.Exec("DELETE FROM badminton_games")
}

func ptrString(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}

func valueOrNilString(p *string) interface{} {
	if p == nil {
		return nil
	}
	return *p
}

func valueOrNilInt(p *int) interface{} {
	if p == nil {
		return nil
	}
	return *p
}