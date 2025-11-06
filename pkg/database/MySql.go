package database

import (
	"log"
	"task-tracker/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySql(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database: " + err.Error())
	}
	return db, nil
}
