package internal

import (
	"task-tracker/internal/Auth"
	"task-tracker/internal/Board"
	"task-tracker/internal/Task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	Task.Module,
	Board.Module,
	Auth.Module,
)
