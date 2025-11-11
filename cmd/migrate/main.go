package main

import (
	"github.com/Nebuska/task-tracker/config"
	"github.com/Nebuska/task-tracker/internal/aAuth"
	"github.com/Nebuska/task-tracker/internal/aBoard"
	"github.com/Nebuska/task-tracker/internal/aBoardUser"
	"github.com/Nebuska/task-tracker/internal/aTask"
	"github.com/Nebuska/task-tracker/internal/aUser"
	"github.com/Nebuska/task-tracker/pkg/database"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error loading .env file " + err.Error())
	}
	db, err := database.NewMySql(nil, cfg, nil)
	if err != nil {
		log.Fatal("Error starting database connection " + err.Error())
	}
	err = db.AutoMigrate(
		&aAuth.UserCredentials{},
		&aUser.User{},
		&aBoard.Board{},
		&aTask.Task{},
		&aBoardUser.BoardUser{},
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
	password, _ := bcrypt.GenerateFromPassword([]byte("Test12345"), bcrypt.DefaultCost)
	Creds := []aAuth.UserCredentials{
		{
			Username: "TEST1",
			Password: string(password),
			User: aUser.User{
				FirstName: "Test 1",
				LastName:  "User",
				Email:     "test@test.test",
			},
		},
		{
			Username: "TEST2",
			Password: string(password),
			User: aUser.User{
				FirstName: "Test 2",
				LastName:  "User",
				Email:     "test@test.test",
			},
		},
	}
	db.Create(&Creds)
	Boards := []aBoard.Board{
		{
			Name: "TEST 1 Board",
		},
		{
			Name: "TEST 2 Board",
		},
	}
	db.Create(&Boards)
	Tasks := []aTask.Task{
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
	BoardUsers := []aBoardUser.BoardUser{
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
