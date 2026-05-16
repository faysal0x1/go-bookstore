package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewBookRepository),
	fx.Provide(NewAuthRepository),
)
