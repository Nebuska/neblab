package main

import (
	"log"

	"github.com/Nebuska/neblab/shared/config"
	"github.com/Nebuska/neblab/shared/database"
	"github.com/Nebuska/neblab/shared/database/mysql"
	"github.com/Nebuska/neblab/tasker/internal/board"
	"github.com/Nebuska/neblab/tasker/internal/boardUser"
	"github.com/Nebuska/neblab/tasker/internal/task"

	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig(database.NewConfig())
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
	db, err := mysql.NewMySql(nil, cfg.DatabaseConfig, nil)
	if err != nil {
		log.Fatal("Error starting database connection " + err.Error())
	}
	err = db.AutoMigrate(
		&board.Board{},
		&task.Task{},
		&boardUser.BoardUser{},
	)
	if err != nil {
		log.Fatal("Error starting database migration " + err.Error())
	}
	log.Println("Database migration succeeded")
	log.Println("Database seeding started")
	Seed(db)
	log.Println("Database seeding succeeded")
}

func Seed(db *gorm.DB) {
	Boards := []board.Board{
		{
			Name: "TEST 1 Board",
		},
		{
			Name: "TEST 2 Board",
		},
	}
	db.Create(&Boards)
	Tasks := []task.Task{
		{
			Name:        "Test Task 1",
			Description: "Test task description",
			BoardID:     Boards[0].ID,
		},
		{
			Name:        "Test Task 2",
			Description: "Test task description",
			BoardID:     Boards[0].ID,
		},
		{
			Name:        "Test Task 3",
			Description: "Test task description",
			BoardID:     Boards[1].ID,
		},
		{
			Name:        "Test Task 4",
			Description: "Test task description",
			BoardID:     Boards[1].ID,
		},
	}
	db.Create(&Tasks)
	BoardUsers := []boardUser.BoardUser{
		{
			UserID:  1,
			BoardID: Boards[0].ID,
			Role:    "",
		},
		{
			UserID:  2,
			BoardID: Boards[0].ID,
			Role:    "",
		},
		{
			UserID:  1,
			BoardID: Boards[1].ID,
			Role:    "",
		},
	}
	db.Create(&BoardUsers)
}
