package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// todo : Use modern libraries like viper or envconfig for environment configuration.
type Config struct {
	DatabaseConnectionString string
	ServerPort               string
	LatestApiVersion         string
	JWTSecret                string
	JWTExpire                time.Duration
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}

	return &Config{
		ServerPort:       os.Getenv("SERVER_PORT"),
		LatestApiVersion: "v1",
		DatabaseConnectionString: os.Getenv("DB_USER") + ":" +
			os.Getenv("DB_PASS") +
			"@tcp(127.0.0.1:3306)/" +
			os.Getenv("DB_NAME") +
			"?charset=utf8mb4&parseTime=True&loc=Local",
		JWTSecret: "THISISTHESECRET",
		JWTExpire: time.Hour,
	}, nil
}
