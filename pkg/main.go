package main

import (
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// AppVersion incrementing
	AppVersion = 0.4
	// AppName names the app
	AppName = "StudRZ"
)

func main() {
	c := LoadConfig()
	db := ConnectDatabase(
		c.DBUser,
		c.DBPassword,
		c.Database,
		c.DBHost,
		c.DBPort,
		c.DBSSl,
		c.TimeZone,
	)
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
		middleware.BasicAuth(
			DatabaseBasicAuth(db),
		),
	)

	// Add tracing
	j := jaegertracing.New(e, nil)
	defer j.Close()
	// Add /metrics Endpoint
	p := prometheus.NewPrometheus(AppName, nil)
	p.Use(e)

	SetAdminUser(db, c)
	// Routes
	RegisterRoutes(
		BundleRoutes(
			// Welcome Page
			CreateRoute("/", e.GET, Welcome(db)),
			CreateRoute("/api", e.GET, Welcome(db)),
			// CRD REST API
			// User Routes
			CreateRoute("/api/user", e.GET, ListUsers(db)),
			CreateRoute("/api/user", e.POST, CreateUser(db)),
			CreateRoute("/api/user/:id", e.GET, GetUser(db)),
			CreateRoute("/api/user/:id", e.DELETE, DeleteUser(db)),
			CreateRoute("/api/user/:id/result", e.GET, GetQRByUser(db)),
			// QuestionairResult Routes
			CreateRoute("/api/result", e.GET, ListQuestionairResult(db)),
			CreateRoute("/api/result", e.POST, CreateQuestionairResult(db)),
			CreateRoute("/api/result/:id", e.GET, GetQuestionairResult(db)),
			CreateRoute("/api/result/:id", e.DELETE, DeleteQuestionairResult(db)),
			CreateRoute("/api/result/:id/consume", e.GET, GetConsumeByQuestionairResult(db)),
			// Consume Routes
			CreateRoute("/api/consume", e.GET, ListConsumes(db)),
			CreateRoute("/api/consume", e.POST, CreateConsume(db)),
			CreateRoute("/api/consume/:id", e.GET, GetConsume(db)),
			CreateRoute("/api/consume/:id", e.DELETE, DeleteConsume(db)),
			// Statistics
			CreateRoute("/api/statistic", e.GET, ListStatistic(db)),
			CreateRoute("/api/statistic/:function/:name", e.GET, CreateStatisticByName(db)),
		))
	// Start server
	e.Logger.Fatal(
		e.Start(":42069"))
}
