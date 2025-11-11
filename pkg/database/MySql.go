package database

import (
	"context"
	"log"
	"task-tracker/config"
	"task-tracker/pkg/logger"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySql(lc fx.Lifecycle, cfg *config.Config, logger *logger.GormLogger) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseConnectionString), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		log.Fatal("Error while connecting to database: " + err.Error())
	}
	if lc != nil {
		lc.Append(fx.Hook{
			/* if I decide to add operations before the program starts
			OnStart: func(ctx context.Context) error {
				return nil
			},*/
			OnStop: func(ctx context.Context) error {
				sql, _ := db.DB()
				return sql.Close()
			},
		})
	}

	return db, nil
}
