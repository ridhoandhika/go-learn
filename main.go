package main

import (
	"santrikoding/backend-api/internal/api"
	"santrikoding/backend-api/internal/component"
	"santrikoding/backend-api/internal/config"
	"santrikoding/backend-api/internal/middleware"
	"santrikoding/backend-api/internal/repository"
	"santrikoding/backend-api/internal/service"

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
