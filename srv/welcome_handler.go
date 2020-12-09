package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Welcome page
func Welcome(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(echo.HeaderAccessControlAllowOrigin, "*")
		return c.JSON(
			http.StatusOK,
			map[string]interface{}{
				"name":    AppName,
				"version": AppVersion,
				"routes": map[string]interface{}{
					"/api/user":        "GET",
					"/api/user/:id":    []string{"GET", "POST", "DELETE"},
					"/api/result":      "GET",
					"/api/result/:id":  []string{"GET", "POST", "DELETE"},
					"/api/consume":     "GET",
					"/api/consume/:id": []string{"GET", "POST", "DELETE"},
				},
				"time": time.Now(),
			},
		)
	}
}
