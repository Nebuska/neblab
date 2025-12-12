package jwtAuth

import "go.uber.org/fx"

var Module = fx.Provide(NewJWTManager)
