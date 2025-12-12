package auth

import (
	"github.com/Nebuska/neblab/account/internal/credentials"
	"github.com/Nebuska/neblab/account/internal/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(credentials.NewCredentialsRepository),
	fx.Provide(user.NewUserRepository),
	fx.Provide(NewAuthService),
)
