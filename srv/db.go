package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// ConnectDatabase returns a dataabse session
func ConnectDatabase(user string, password string, database string, host string, port int, ssl bool, timezone string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf(
			"user=%s password=%s database=%s host=%s port=%d sslmode=%s TimeZone=%s",
			user,
			password,
			database,
			host,
			port,
			func(s bool) string {
				switch(s) {
				case true:
					return "enable"
				case false:
					return "disable"
				default:
					return "disable"
				}
			}(ssl),
			timezone,
		),
	), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	if err := db.AutoMigrate(
		&QuestionairResult{},
		&User{},
		&Consume{},
		&Statistic{},
	); err != nil {
		panic(err)
	}
	return db
}
