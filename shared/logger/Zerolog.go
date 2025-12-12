package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func NewZeroLogger() (*Logger, error) {
	zerolog.TimeFieldFormat = time.RFC3339

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: time.TimeOnly,
	}

	logger := zerolog.New(consoleWriter).
		With().Timestamp().Caller().Logger()

	return &Logger{logger}, nil
}
