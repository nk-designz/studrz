package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	// StatisticFunctions in a map bc im lazy
	StatisticFunctions = map[string]StatisticFunc{
		"mean":      Mean,
		"deviation": Deviation,
	}
)

// ListStatistic returns a statistic
func ListStatistic(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var stList []*Statistic
		if err := db.Table("statistics").Find(&stList).Error; err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Code: http.StatusInternalServerError,
					Message: "Database error",
					Payload: fmt.Sprint(err),
				},
			)
		}
		return c.JSON(http.StatusOK, stList)
	}
}

// CreateStatisticByName creates a statistic result
func CreateStatisticByName(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		function := c.Param("function")
		name := c.Param("name")
		st := &Statistic{
			Key: name,
			Calc: function,
		}
		if err := st.Calculate(db); err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusNoContent,
				&ErrorMessage{
					Code:    http.StatusNoContent,
					Message: "undefined",
					Payload: []string{name, function},
				},
			)
		}
		if err := st.Save(db); err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusNoContent,
				&ErrorMessage{
					Code:    http.StatusInternalServerError,
					Message: "Database error",
					Payload: name,
				},
			)
		}
		return c.JSON(
			http.StatusOK,
			st,
		)
	}
}

// Mean function
func Mean(s *Statistic, qa []QuestionairResult) error {
	if s.Findings <= 0 {
		return fmt.Errorf("Dont devide by 0")
	}
	for _, q := range qa {
		s.Value += float64(q.Score)
	}
	s.Value /= float64(s.Findings)
	return nil
}

// Deviation function
func Deviation(s *Statistic, qa []QuestionairResult) error {
	var mean, sd float64
	if s.Findings <= 0 {
		return fmt.Errorf("Dont devide by 0")
	}
	for _, q := range qa {
		mean += float64(q.Score)
	}
	mean /= float64(s.Findings)

	for j := 0; j < int(s.Findings); j++ {
		sd += math.Pow(float64(qa[j].Score)-mean, 2)
	}
	s.Value = math.Sqrt(sd / float64(s.Findings))
	return nil
}
