package database

import (
	"os"
)

type Config struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

func NewConfig() *Config {
	return &Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
	}
}
