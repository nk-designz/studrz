package main

import (
	"gorm.io/gorm"
)

// StatisticFunc calculate the statistic
type StatisticFunc func(*Statistic, []QuestionairResult) error

// Statistic type
type Statistic struct {
	gorm.Model
	ID       uint64  `json:"id"`
	Calc     string  `json:"calc"`
	Value    float64 `json:"value"`
	Key      string  `json:"key"`
	Findings uint    `json:"findings"`
}

// Calculate the statistic
func (s *Statistic) Calculate(db *gorm.DB) error {
	var qr []QuestionairResult
	if err := db.Table("consumes").Joins(
		"JOIN questionair_results on questionair_results.id = consumes.questionair_result_id",
	).Where("consumes.name = ?", s.Key).Scan(&qr).Error; err != nil {
		return err
	}
	s.Findings = uint(len(qr))
	return StatisticFunctions[s.Calc](s, qr)
}

// Save to Database
func (s *Statistic) Save(db *gorm.DB) error {
	return db.Save(s).Error
}
