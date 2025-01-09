package modules

import (
	"user-management/internal/config"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	options := []fx.Option{
		InternalModule,
		UserModule,
	}

	if !config.IsLocalScope() {
		options = append(options, fx.NopLogger)
	}

	return fx.New(
		fx.Options(options...),
	)
}
