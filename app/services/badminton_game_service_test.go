package services_test

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-go-test/app/models"
	"gin-go-test/app/services"
)

func ptrString(s string) *string { return &s }
func ptrInt(i int) *int          { return &i }
func ptrTime(t time.Time) *time.Time { return &t }

func TestBatchUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect db: %v", err)
	}

	if err := db.AutoMigrate(&models.BadmintonGame{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	service := services.NewBadmintonGameService(db)

	original := models.BadmintonGame{
		Player1:   ptrString("Tom"),
		Player2:   ptrString("Jack"),
		Score1:    ptrInt(21),
		Score2:    ptrInt(15),
		Location:  ptrString("Old Location"),
		MatchTime: ptrTime(time.Now()),
	}
	if err := db.Create(&original).Error; err != nil {
		t.Fatalf("failed to create record: %v", err)
	}

	t.Logf("🧾 Before update: Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v", *original.Player1, *original.Player2, *original.Score1, *original.Score2, *original.Location)

	update := models.BadmintonGame{
		Id:     original.Id,
		Score1: ptrInt(22),
		Score2: ptrInt(20),
		// 其他字段为 nil 表示不更新
	}

	if err := service.BatchUpdate([]models.BadmintonGame{update}); err != nil {
		t.Fatalf("BatchUpdate failed: %v", err)
	}

	var updated models.BadmintonGame
	if err := db.First(&updated, original.Id).Error; err != nil {
		t.Fatalf("failed to query updated record: %v", err)
	}
	t.Logf("✅ After update: Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v", *updated.Player1, *updated.Player2, *updated.Score1, *updated.Score2, *updated.Location)

	if updated.Player1 == nil || *updated.Player1 != "Tom" {
		t.Errorf("Player1 field changed unexpectedly: got %v", updated.Player1)
	}
	if updated.Player2 == nil || *updated.Player2 != "Jack" {
		t.Errorf("Player2 field changed unexpectedly: got %v", updated.Player2)
	}
	if updated.Location == nil || *updated.Location != "Old Location" {
		t.Errorf("Location field changed unexpectedly: got %v", updated.Location)
	}

	if updated.Score1 == nil || *updated.Score1 != 22 {
		t.Errorf("Score1 not updated correctly: got %v", updated.Score1)
	}
	if updated.Score2 == nil || *updated.Score2 != 20 {
		t.Errorf("Score2 not updated correctly: got %v", updated.Score2)
	}
}