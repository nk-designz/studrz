package main

import (
	"github.com/kelseyhightower/envconfig"
)

// Config for this app
type Config struct {
	AdminPassword string
	DBPassword    string
	DBUser        string
	Database      string
	TimeZone      string
	DBPort        int
	DBHost        string
	DBSSl         bool
}

// LoadConfig from env
func LoadConfig() Config {
	var c Config
	err := envconfig.Process(AppName, &c)
	if err != nil {
		panic(err)
	}
	return c
}
