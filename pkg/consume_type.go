package main

import "gorm.io/gorm"

// Consume results
type Consume struct {
	gorm.Model          `json:"-"`
	ID                  uint64 `json:"id,omitempty" gorm:"primaryKey"`
	QuestionairResultID uint64 `json:"questionair_result_id"`
	QuestionairResult   QuestionairResult `json:"-"`
	Name                string `json:"name"`
	Rate                uint8  `json:"rate"`
}
