package main

import (
	"log"
	"task-tracker/config"
	"task-tracker/internal/Task"
	"task-tracker/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
	db, err := database.NewMySql(nil, cfg)
	if err != nil {
		log.Fatal("Error starting database connection " + err.Error())
	}
	err = db.AutoMigrate(
		&Task.Task{},
	)
	if err != nil {
		log.Fatal("Error starting database migration " + err.Error())
	}
	log.Println("Database migration succeeded")
}
