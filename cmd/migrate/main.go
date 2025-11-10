package main

import (
	"log"
	"task-tracker/config"
	"task-tracker/internal/Auth"
	"task-tracker/internal/Board"
	"task-tracker/internal/BoardUser"
	"task-tracker/internal/Task"
	"task-tracker/internal/User"
	"task-tracker/pkg/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		&Auth.UserCredentials{},
		&User.User{},
		&Board.Board{},
		&Task.Task{},
		&BoardUser.BoardUser{},
	)
	if err != nil {
		log.Fatal("Error starting database migration " + err.Error())
	}
	log.Println("Database migration succeeded")
	log.Println("Database seeding started")
	Seed(db, cfg)
	log.Println("Database seeding succeeded")
}

func Seed(db *gorm.DB, cfg *config.Config) {
	password, _ := bcrypt.GenerateFromPassword([]byte("Test12345"), bcrypt.DefaultCost)
	Creds := []Auth.UserCredentials{
		{
			Username: "TEST1",
			Password: string(password),
			User: User.User{
				FirstName: "Test 1",
				LastName:  "User",
				Email:     "test@test.test",
			},
		},
		{
			Username: "TEST2",
			Password: string(password),
			User: User.User{
				FirstName: "Test 2",
				LastName:  "User",
				Email:     "test@test.test",
			},
		},
	}
	db.Create(&Creds)
	Boards := []Board.Board{
		{
			Name: "TEST 1 Board",
		},
		{
			Name: "TEST 2 Board",
		},
	}
	db.Create(&Boards)
	Tasks := []Task.Task{
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
	BoardUsers := []BoardUser.BoardUser{
		{
			UserID:  Creds[0].ID,
			BoardID: Boards[0].ID,
			Role:    "",
		},
		{
			UserID:  Creds[1].ID,
			BoardID: Boards[0].ID,
			Role:    "",
		},
		{
			UserID:  Creds[0].ID,
			BoardID: Boards[1].ID,
			Role:    "",
		},
	}
	db.Create(&BoardUsers)
}
