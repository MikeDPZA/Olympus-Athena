package main

import (
	_ "Olympus-Athena/cmd/docs"
	"Olympus-Athena/pkg/database"
	"Olympus-Athena/pkg/handlers"
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
	fiberApp  *fiber.App
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

func (a *AthenaApp) RegisterFiber() {
	a.fiberApp = fiber.New(fiber.Config{
		AppName: "Athena API",
	})

	a.fiberApp.Use(recover.New())
	a.fiberApp.Use(cors.New())
	a.fiberApp.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/athena/docs": "/athena/docs/index.html",
			"/athena":      "/athena/docs/index.html",
		},
	}))
	a.fiberApp.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	a.fiberApp.Get("/athena/docs/*", swagger.FiberWrapHandler())

	athena := a.fiberApp.Group("/athena")
	{
		athena.Get("/health", handlers.HealthHandler)
	}
}

// @title Athena API
// @version 1.0
// @description This service manages Identity and Access Management
func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("No environment files found")
	}

	dbConnString, _ := os.LookupEnv("DbConnectionString")

	app := NewAthenaApp(AthenaAppConfig{
		DbConnectionString: dbConnString,
	})

	app.RegisterFiber()
	app.RegisterDb()
	defer app.db.Disconnect()

	slog.Error("Critical Error launching API", "Error", app.fiberApp.Listen(":4200"))
}
