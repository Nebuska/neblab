package config

import (
	"os"
	"time"

	"github.com/Nebuska/neblab/shared/database"
)

type Config struct {
	DatabaseConfig   *database.Config
	ServerPort       string
	LatestApiVersion string
	JWTSecret        string
	JWTExpire        time.Duration
}

func NewConfig(dbConfig *database.Config) (*Config, error) {
	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRES"))
	if err != nil {
		return nil, err
	}
	return &Config{
		DatabaseConfig:   dbConfig,
		ServerPort:       os.Getenv("PORT"),
		LatestApiVersion: "v1",
		JWTSecret:        os.Getenv("JWT_SECRET"),
		JWTExpire:        duration,
	}, nil
}
