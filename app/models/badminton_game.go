package models

import (
	"time"
)

// BadmintonGame 数据模型  Key
type BadmintonGame struct {
	Id        *int       `gorm:"column:id" json:"id" validate:"max=255"`
	Player1   *string    `gorm:"column:player1" json:"player1" validate:"required,max=255"`
	Player2   *string    `gorm:"column:player2" json:"player2" validate:"required,max=255"`
	Score1    *int       `gorm:"column:score1" json:"score1" validate:"max=255"`
	Score2    *int       `gorm:"column:score2" json:"score2" validate:"max=255"`
	Location  *string    `gorm:"column:location" json:"location" validate:"max=255"`
	MatchTime *time.Time `gorm:"column:match_time" json:"match_time" validate:"max=255"`
}
