package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConnectionString string
	ServerPort               string
	LatestApiVersion         string
	JWTSecret                string
	JWTExpire                time.Duration
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}

	duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		return nil, err
	}
	return &Config{
		ServerPort:       os.Getenv("SERVER_PORT"),
		LatestApiVersion: "v1",
		DatabaseConnectionString: os.Getenv("DB_USER") + ":" +
			os.Getenv("DB_PASS") + "@tcp(" +
			os.Getenv("DB_HOST") + ":" +
			os.Getenv("DB_PORT") + ")/" +
			os.Getenv("DB_NAME") +
			"?charset=utf8mb4&parseTime=True&loc=Local",
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTExpire: duration,
	}, nil
}
