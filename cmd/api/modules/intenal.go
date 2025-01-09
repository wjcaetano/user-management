package modules

import (
	"user-management/app/user/entrypoint/rest/handler"
	"user-management/internal/config"
	"user-management/internal/db"
	"user-management/internal/router"
	"user-management/internal/routes"
	"user-management/internal/server"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

var internalFactory = fx.Provide(
	config.NewConfig,
	db.NewDatabase,
	router.NewRouter,
)

var InternalModule = fx.Options(
	internalFactory,
	fx.Invoke(
		server.StartHTTPServer,
		func(router *chi.Mux, handler *handler.Handler) {
			routes.RegisterRoutes(router, handler)
		},
	),
)
