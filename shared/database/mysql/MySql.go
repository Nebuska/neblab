package mysql

import (
	"context"
	"fmt"

	"github.com/Nebuska/neblab/shared/database"
	"github.com/Nebuska/neblab/shared/logger"

	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySql(lc fx.Lifecycle, cfg *database.Config, logger *logger.GormLogger) (*gorm.DB, error) {
	DSN := newConnectionString(cfg)
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}
	if lc != nil {
		lc.Append(fx.Hook{
			/* if I decide to add operations before the program starts
			OnStart: func(ctx context.Context) error {
				return nil
			},*/
			OnStop: func(ctx context.Context) error {
				sql, err := db.DB()
				if err != nil {
					return err
				}
				return sql.Close()
			},
		})
	}

	return db, nil
}

func newConnectionString(cfg *database.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		"charset=utf8mb4&parseTime=True&loc=Local")
}
