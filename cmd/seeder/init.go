package main

import (
	"database/sql"

	"github.com/3tagger/echo-sample-arch/internal/seeder"
	userrepository "github.com/3tagger/echo-sample-arch/internal/user/repository"
)

func initSeederMap(db *sql.DB) map[string]seeder.SeederExecutor {
	userRepository := userrepository.NewUserRepositoryPostgreSQL(db)
	userService := seeder.NewuserSeederService(userRepository)
	userSeederExec := seeder.NewSeederExecutor(userService)

	return map[string]seeder.SeederExecutor{
		userEntity: userSeederExec,
	}
}
