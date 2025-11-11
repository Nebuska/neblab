package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

type zeroLogger struct {
	Zero *zerolog.Logger
}

func NewZeroLogger() (*Logger, error) {
	zerolog.TimeFieldFormat = time.RFC3339

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    false,
		TimeFormat: time.TimeOnly,
	}

	fileWriter, err := os.OpenFile("logs.jsonl", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, fileWriter)

	logger := zerolog.New(multi).
		With().Timestamp().Caller().Logger()

	return &Logger{logger}, err
}
