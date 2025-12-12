package postgres

import (
	"context"
	"fmt"

	"github.com/Nebuska/neblab/shared/database"
	"github.com/Nebuska/neblab/shared/logger"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(lc fx.Lifecycle, cfg *database.Config, logger *logger.GormLogger) (*gorm.DB, error) {
	dsn := newConnectionString(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}
	if lc != nil {
		lc.Append(fx.Hook{
			/* Might be used on migrations
			OnStart: func(ctx context.Context) error {
				return nil
			},*/
			OnStop: func(ctx context.Context) error {
				s, err := db.DB()
				if err != nil {
					return err
				}
				return s.Close()
			},
		})
	}

	return db, nil
}

func newConnectionString(cfg *database.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Database,
		cfg.Port,
		"disable")
}
