package main

import (
	"ridhoandhika/backend-api/internal/api"
	"ridhoandhika/backend-api/internal/component"
	"ridhoandhika/backend-api/internal/config"
	"ridhoandhika/backend-api/internal/middleware"
	"ridhoandhika/backend-api/internal/repository"
	"ridhoandhika/backend-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRepository, cacheConnection)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()

	api.NewAuth(app, userService, authMiddleware)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
