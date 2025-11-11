package api

import (
	v1 "github.com/Nebuska/task-tracker/api/v1"

	"go.uber.org/fx"
)

var Module = fx.Options(
	v1.Module)
