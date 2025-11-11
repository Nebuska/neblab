package logger

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

type GormLogger struct {
	logger *Logger
}

func NewGormLogger(logger *Logger) *GormLogger {
	l := logger.With().Str("service", "GORM").Logger()
	return &GormLogger{logger: &Logger{l}}
}

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g GormLogger) Info(ctx context.Context, s string, i ...interface{}) {
	g.logger.Info().Msg(fmt.Sprintf(s, i...))
}

func (g GormLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	g.logger.Warn().Msg(fmt.Sprintf(s, i...))
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	g.logger.Error().Msg(fmt.Sprintf(s, i...))
}

func (g GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	g.logger.Trace().Str("sql", sql).Int64("rows", rows).Err(err).Msg("SQL executed")
}
