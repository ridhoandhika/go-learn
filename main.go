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
	// cacheConnection := component.GetCacheConnection()

	userRepository := repository.User(dbConnection)
	personalInfoRepository := repository.PersonalInformation((dbConnection))

	userService := service.User(userRepository)
	personalInfoService := service.PersonalInformation(personalInfoRepository)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()
	apiRoutes := app.Group("api")
	// route
	api.Auth(apiRoutes.(*fiber.Group), userService, authMiddleware)
	api.PersonalInformation(apiRoutes.(*fiber.Group), personalInfoService, authMiddleware)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
