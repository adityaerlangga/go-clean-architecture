package main

import (
	"go-clean-architecture/app"
	"go-clean-architecture/app/database"
	"go-clean-architecture/services/logics"
	"go-clean-architecture/services/repository"
)

func main() {
	db := database.SetupDatabase()

	UserRepository := repository.InitUserRepository(db)
	UserLogic := logics.InitUserLogic(UserRepository)

	server := app.SetupRoutes(UserLogic)
	server.Run(":8081")
}
