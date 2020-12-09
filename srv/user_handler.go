package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ListUsers get a specific user
func ListUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, "*")
		var users []*User
		if err := db.Table("users").Order("id ASC").Scan(&users).Error; err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Code: http.StatusInternalServerError,
					Message: "Database error",
					Payload: err,
				},
			)
		}
		for _, k := range users {
			k.Password = ""
		}
		return c.JSON(http.StatusOK, users)
	}
}

// GetUser get a specific user
func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, "*")
		id := c.Param("id")
		var user User
		if err := db.Table("users").Where("id = ?", id).Scan(&user).Error; err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				&ErrorMessage{
					Code: http.StatusInternalServerError,
					Message: "Dataabse error",
					Payload: id,
				},
			)
		}
		user.Password = ""
		return c.JSON(http.StatusOK, user)
	}
}

// CreateUser create a user
func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, "*")
		user := new(User)
		if err := c.Bind(user); err != nil {
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: user,
					Message: "Cannot read data",
					Code:    http.StatusBadRequest,
				},
			)
		}
		if user.Password == "" || user.Username == "" || user.Role > 3 {
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: user,
					Message: "Not enough data",
					Code:    http.StatusBadRequest,
				},
			)
		}
		c.Logger().Info(
			fmt.Sprintf("Creating %dth user %s with role %d", user.ID, user.Username, user.Role))
		hashedUser, err := user.GetHashed()
		if err != nil {
			return c.JSON(
				http.StatusBadRequest,
				&ErrorMessage{
					Payload: user,
					Message: "Password not hashable",
					Code:    http.StatusInternalServerError,
				},
			)
		}
		db.Create(hashedUser)
		user.Password = ""
		return c.JSON(http.StatusCreated, user)
	}
}

// DeleteUser deletes a user
func DeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, "*")
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
		user := new(User)
		db.Table("users").Where("id = ?", id).Scan(&user)
		db.Delete(&user)
		user.Password = ""
		return c.JSON(http.StatusOK, user)
	}
}
