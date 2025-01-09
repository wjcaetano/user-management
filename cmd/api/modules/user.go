package modules

import (
	"user-management/app/user/entrypoint/rest/handler"
	"user-management/app/user/repository"
	repositorySQL "user-management/app/user/repository/sql"
	"user-management/app/user/service"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var userFactory = fx.Options(
	fx.Provide(
		func(db *gorm.DB) repository.UserRepository {
			return repositorySQL.NewUserRepository(db)
		},
	),

	fx.Provide(
		func(repo repository.UserRepository) service.UserService {
			return service.NewUserService(repo)
		},
	),

	fx.Provide(
		handler.NewHandler,
	),
)

var UserModule = fx.Options(
	userFactory,
)
