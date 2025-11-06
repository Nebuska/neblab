package Task

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewTaskService),
	fx.Provide(NewTaskRepository),
)
