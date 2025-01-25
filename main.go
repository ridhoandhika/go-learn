package main

import (
	"ridhoandhika/backend-api/internal/api"
	"ridhoandhika/backend-api/internal/component"
	"ridhoandhika/backend-api/internal/config"
	"ridhoandhika/backend-api/internal/middleware"
	"ridhoandhika/backend-api/internal/repository"
	"ridhoandhika/backend-api/internal/service"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	// cacheConnection := component.GetCacheConnection()

	userRepository := repository.User(dbConnection)
	personalInfoRepository := repository.PersonalInformation((dbConnection))
	workExperienceRepository := repository.WorkExperience((dbConnection))
	educationRepository := repository.Education((dbConnection))

	userService := service.User(userRepository)
	personalInfoService := service.PersonalInformation(personalInfoRepository)
	workExperienceService := service.WorkExperience(workExperienceRepository)
	educationService := service.Education(educationRepository)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()
	// Tentukan konfigurasi Swagger
	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/swagger",
		Title:    "Swagger API Docs",
	}

	// Gunakan middleware Swagger
	app.Use(swagger.New(cfg))

	apiRoutes := app.Group("api")
	// route

	api.Auth(apiRoutes.(*fiber.Group), userService, authMiddleware)
	api.PersonalInformation(apiRoutes.(*fiber.Group), personalInfoService, authMiddleware)
	api.WorkExperience(apiRoutes.(*fiber.Group), workExperienceService, authMiddleware)
	api.Education(apiRoutes.(*fiber.Group), educationService, authMiddleware)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
