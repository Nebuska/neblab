package internal

import (
	"github.com/Nebuska/task-tracker/internal/auth"
	"github.com/Nebuska/task-tracker/internal/board"
	"github.com/Nebuska/task-tracker/internal/task"

	"go.uber.org/fx"
)

var Module = fx.Options(
	task.Module,
	board.Module,
	auth.Module,
)
