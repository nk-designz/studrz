package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ListConsumes returns a list of Consumes
func ListConsumes(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var consumeList []*Consume
		if err := db.Table("consumes").Order("id ASC").Scan(&consumeList).Error; err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Payload: err,
					Message: "Database error",
					Code:    http.StatusInternalServerError,
				},
			)
		}
		return c.JSON(http.StatusOK, consumeList)
	}
}

// GetConsume returns a specific Consume
func GetConsume(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: id,
					Message: "Could not parse to number",
					Code:    http.StatusBadRequest,
				},
			)
		}
		cr := new(Consume)
		if err := db.Table("consumes").Where("id = ?", id).Scan(&cr).Error; err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Payload: err,
					Message: "Database error",
					Code:    http.StatusInternalServerError,
				},
			)
		}
		return c.JSON(http.StatusOK, cr)
	}
}

// CreateConsume creates consume
func CreateConsume(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		cr := new(Consume)
		if err := c.Bind(&cr); err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: cr,
					Message: "Cannot read data",
					Code:    http.StatusBadRequest,
				},
			)
		}
		if err := db.Create(cr).Error; err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: cr,
					Message: "Database error",
					Code:    http.StatusInternalServerError,
				},
			)
		}
		return c.JSON(http.StatusOK, cr)
	}
}

// DeleteConsume deletes a specific consume
func DeleteConsume(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: id,
					Message: "Could not parse to number",
					Code:    http.StatusBadRequest,
				},
			)
		}
		cr := new(Consume)
		if err := db.Table("consumes").Where("id = ?", id).Scan(&cr); err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: id,
					Message: "No user with this ID exists",
					Code:    http.StatusBadRequest,
				},
			)
		}
		if err := db.Delete(&cr).Error; err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: id,
					Message: "Database error",
					Code:    http.StatusInternalServerError,
				},
			)
		}
		return c.JSON(http.StatusOK, cr)
	}
}

// GetConsumeByQuestionairResult returns all consumes of a qr
func GetConsumeByQuestionairResult(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cList []Consume
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: id,
					Message: "Could not parse to number",
					Code:    http.StatusBadRequest,
				},
			)
		}
		if err := db.Table("consumes").Where("questionair_result_id = ?", id).Scan(&cList).Error; err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Code:    http.StatusInternalServerError,
					Message: "Not found by id",
					Payload: id,
				},
			)
		}
		return c.JSON(http.StatusOK, cList)
	}
}
