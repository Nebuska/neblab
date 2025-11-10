package v1

import (
	"task-tracker/api/v1/base"
	"task-tracker/api/v1/board"
	"task-tracker/api/v1/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	task.Module,
	board.Module,
	base.Module,
)
