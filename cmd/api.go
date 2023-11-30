package main

import (
	_ "Olympus-Athena/cmd/docs"
	"Olympus-Athena/pkg/controllers"
	"Olympus-Athena/pkg/database"
	"Olympus-Athena/pkg/repositories"
	"Olympus-Athena/pkg/routes"
	"Olympus-Athena/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/joho/godotenv"
	swagger "github.com/swaggo/fiber-swagger"
	"log/slog"
	"os"
)

type AthenaAppConfig struct {
	DbConnectionString string
}

type AthenaApp struct {
	appConfig AthenaAppConfig
	db        database.AthenaDb
}

func (a *AthenaApp) RegisterDb() {
	db := database.NewAthenaMySqlDb(database.AthenaDbConfig{
		ConnectionString: a.appConfig.DbConnectionString,
	})
	db.Connect()
	a.db = db
}

func NewAthenaApp(appConfig AthenaAppConfig) *AthenaApp {
	return &AthenaApp{
		appConfig: appConfig,
	}
}

// @title Athena API
// @version 1.0
// @description This service manages Identity and Access Management
// @Produce json
// @Accept json
func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("No environment files found")
	}

	dbConnString, _ := os.LookupEnv("DbConnectionString")

	app := NewAthenaApp(AthenaAppConfig{
		DbConnectionString: dbConnString,
	})

	// Dependencies
	app.RegisterDb()
	defer app.db.Disconnect()

	// repos
	policyRepo := repositories.NewPolicyMySqlRepository(&app.db)

	// services
	policyService := services.NewPolicyMySqlService(policyRepo)

	// Fiber
	fiberApp := fiber.New(fiber.Config{
		AppName: "Athena API",
	})

	fiberApp.Use(recover.New())
	fiberApp.Use(cors.New())
	fiberApp.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/athena/docs": "/athena/docs/index.html",
			"/athena":      "/athena/docs/index.html",
		},
	}))
	fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	fiberApp.Get("/athena/docs/*", swagger.FiberWrapHandler())

	athena := fiberApp.Group("/athena")
	{
		athena.Get("/health", controllers.HealthHandler)

		v1 := athena.Group("/v1")
		{
			routes.SetupPolicyRoutes(v1, controllers.NewPolicyController(policyService))
			routes.SetupOAuthRoutes(v1, controllers.NewOAuthClientController())
		}

	}

	slog.Error("Critical Error launching API", "Error", fiberApp.Listen(":4200"))
}
