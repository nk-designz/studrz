package main

import (
	"time"

	"gorm.io/gorm"
)

// https://gorm.io/docs/belongs_to.html

// QuestionairResult for a test
type QuestionairResult struct {
	gorm.Model  `json:"-"`
	ID          uint64    `json:"id"  gorm:"primaryKey"`
	UserID      uint64    `json:"user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	Score       uint8     `json:"score" gorm:"index"`
	FillingDate time.Time `json:"filling_date"`
}
