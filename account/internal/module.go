package internal

import (
	"github.com/Nebuska/neblab/account/internal/auth"
	"github.com/Nebuska/neblab/account/internal/credentials"
	"github.com/Nebuska/neblab/account/internal/session"
	"github.com/Nebuska/neblab/account/internal/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(credentials.NewCredentialsRepository),
	fx.Provide(user.NewUserRepository),
	fx.Provide(session.NewRepository),
	fx.Provide(auth.NewAuthService),
)
