package main

import (
	"log"
	"task-tracker/config"
	"task-tracker/internal/Board"
	"task-tracker/internal/BoardUser"
	"task-tracker/internal/Task"
	"task-tracker/internal/User"
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
		&User.User{},
		&Board.Board{},
		&Task.Task{},
		&BoardUser.BoardUser{},
	)
	if err != nil {
		log.Fatal("Error starting database migration " + err.Error())
	}
	log.Println("Database migration succeeded")
}
