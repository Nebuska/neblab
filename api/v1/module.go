package v1

import (
	"github.com/Nebuska/task-tracker/api/v1/base"
	"github.com/Nebuska/task-tracker/api/v1/board"
	"github.com/Nebuska/task-tracker/api/v1/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	task.Module,
	board.Module,
	base.Module,
)
