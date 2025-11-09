package Auth

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewAuthRepository),
	fx.Provide(NewAuthService),
)
