package api

import (
	"github.com/Nebuska/neblab/tasker/api/v1"

	"go.uber.org/fx"
)

var Module = fx.Options(
	v1.Module)
