package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ListQuestionairResult returns a list of QRs
func ListQuestionairResult(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, c.Request().Header.Get(echo.HeaderOrigin))
		var qrList []*QuestionairResult
		db.Table("questionair_results").Where("deleted_at IS NULL").Order("id ASC").Scan(&qrList)
		return c.JSON(http.StatusOK, qrList)
	}
}

// GetQuestionairResult returns a specific QRs
func GetQuestionairResult(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, c.Request().Header.Get(echo.HeaderOrigin))
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
		qr := new(QuestionairResult)
		db.Table("questionair_results").Where("id = ?", id).Scan(&qr)
		return c.JSON(http.StatusOK, qr)
	}
}

// CreateQuestionairResult what do u think
func CreateQuestionairResult(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, c.Request().Header.Get(echo.HeaderOrigin))
		qr := new(QuestionairResult)
		if err := c.Bind(&qr); err != nil {
			fmt.Println(err)
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: qr,
					Message: "Cannot read data",
					Code:    http.StatusBadRequest,
				},
			)
		}
		db.Create(qr)
		return c.JSON(http.StatusOK, qr)
	}
}

// DeleteQuestionairResult deletes a specific QR
func DeleteQuestionairResult(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, c.Request().Header.Get(echo.HeaderOrigin))
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
		qr := new(QuestionairResult)
		db.Table("questionair_results").Where("id = ?", id).Scan(&qr)
		db.Delete(&qr)
		return c.JSON(http.StatusOK, qr)
	}
}

// GetQRByUser returns all qr of a user
func GetQRByUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, c.Request().Header.Get(echo.HeaderOrigin))
		var qList []QuestionairResult
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
		if err := db.Table("questionair_results").Where("user_id = ?", id).Scan(&qList).Error; err != nil {
			c.Logger().Warn(err)
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Code: http.StatusInternalServerError,
					Message: "Not found by id",
					Payload: id,
				},
			)
		}
		return c.JSON(http.StatusOK, qList)
	}
}