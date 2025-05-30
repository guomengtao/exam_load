package biz_test

import (
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gin-go-test/app/models"
	"gin-go-test/app/biz"
	"gin-go-test/app/services"  // 导入 services 包
)

func ptrString(s string) *string     { return &s }
func ptrInt(i int) *int              { return &i }
func ptrTime(t time.Time) *time.Time { return &t }

func TestBizBatchUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect db: %v", err)
	}
	if err := db.AutoMigrate(&models.BadmintonGame{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	// 先创建 service 实例
	svc := services.NewBadmintonGameService(db)
	// 传入 service 实例创建 biz 实例
	bizLayer := biz.NewBadmintonGameBiz(svc)

	original := models.BadmintonGame{
		Player1:   ptrString("Tom"),
		Player2:   ptrString("Jack"),
		Score1:    ptrInt(21),
		Score2:    ptrInt(15),
		Location:  ptrString("Old Gym"),
		MatchTime: ptrTime(time.Now()),
	}

	if err := db.Create(&original).Error; err != nil {
		t.Fatalf("failed to create: %v", err)
	}

	t.Logf("🧾 Before update: Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v",
		*original.Player1, *original.Player2, *original.Score1, *original.Score2, *original.Location)

	update := models.BadmintonGame{
		Id:     original.Id,
		Score1: ptrInt(25),
		Score2: ptrInt(23),
	}

	// 这里传入指针切片
	if err := bizLayer.BatchUpdate([]*models.BadmintonGame{&update}); err != nil {
		t.Fatalf("BatchUpdate failed: %v", err)
	}

	var updated models.BadmintonGame
	if err := db.First(&updated, *original.Id).Error; err != nil {
		t.Fatalf("query failed: %v", err)
	}

	t.Logf("✅ After update: Player1=%v, Player2=%v, Score1=%v, Score2=%v, Location=%v",
		*updated.Player1, *updated.Player2, *updated.Score1, *updated.Score2, *updated.Location)

	if *updated.Player1 != "Tom" || *updated.Player2 != "Jack" || *updated.Location != "Old Gym" {
		t.Errorf("Unexpected overwrite: %+v", updated)
	}
	if *updated.Score1 != 25 || *updated.Score2 != 23 {
		t.Errorf("Score update failed: %+v", updated)
	}
}