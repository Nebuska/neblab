package internal

import (
	"github.com/Nebuska/neblab/tasker/internal/board"
	"github.com/Nebuska/neblab/tasker/internal/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	task.Module,
	board.Module,
)
