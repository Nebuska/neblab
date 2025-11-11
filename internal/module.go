package internal

import (
	"github.com/Nebuska/task-tracker/internal/aAuth"
	"github.com/Nebuska/task-tracker/internal/aBoard"
	"github.com/Nebuska/task-tracker/internal/aTask"

	"go.uber.org/fx"
)

var Module = fx.Options(
	aTask.Module,
	aBoard.Module,
	aAuth.Module,
)
