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
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	// cacheConnection := component.GetCacheConnection()

	userRepository := repository.User(dbConnection)
	personalInfoRepository := repository.PersonalInformation(dbConnection)
	workExperienceRepository := repository.WorkExperience(dbConnection)
	educationRepository := repository.Education(dbConnection)
	skillRepository := repository.Skill(dbConnection)
	certificationRepository := repository.Certification(dbConnection)

	userService := service.User(userRepository)
	personalInfoService := service.PersonalInformation(personalInfoRepository)
	workExperienceService := service.WorkExperience(workExperienceRepository)
	educationService := service.Education(educationRepository)
	skillService := service.Skill(skillRepository)
	certificationService := service.Certification(certificationRepository)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:8081, http://localhost:8080", // Membolehkan domain tertentu
		AllowMethods: "GET,POST,PUT,DELETE",                                                 // Metode HTTP yang diizinkan
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",                         // Header yang diizinkan
	}))

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${url} ${status} - ${latency} ${bytesSent}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

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
	api.Skill(apiRoutes.(*fiber.Group), skillService, authMiddleware)
	api.Certification(apiRoutes.(*fiber.Group), certificationService, authMiddleware)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
