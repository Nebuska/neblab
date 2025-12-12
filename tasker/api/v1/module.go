package v1

import (
	"github.com/Nebuska/neblab/tasker/api/v1/board"
	"github.com/Nebuska/neblab/tasker/api/v1/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	task.Module,
	board.Module,
)
