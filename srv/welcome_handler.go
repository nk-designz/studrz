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
				"routes": 	map[string]interface{}{
					"/": 						map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of all routes",
					},
					"/api": 					map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of all routes",
					},
					"/api/user": 				map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of all registered users",
					},
					"/api/user/:id": 			map[string]interface{}{
						"methods":		[]string{"GET", "POST", "DELETE"},
						"description": "returns a user by it's id",
					},
					"/api/user/:id/result": 	map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of a results by it's user id",
					},
					"/api/result": 				map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of all results",
					},
					"/api/result/:id": 			map[string]interface{}{
						"methods":     []string{"GET", "POST", "DELETE"},
						"description": "returns a result by it's id",
					},
					"/api/consume":				map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of consumes",
					},
					"/api/consume/:id": 	   	map[string]interface{}{
						"methods":     []string{"GET", "POST", "DELETE"},
						"description": "returns a consume by it's id",
					},
					"/api/result/:id/consume": 	map[string]interface{}{
						"methods":     "GET",
						"description": "returns a list of a consumes by it's result id",
					},
				},
				"time": time.Now(),
			},
		)
	}
}
