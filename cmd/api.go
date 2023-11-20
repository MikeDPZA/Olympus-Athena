package main

import (
	_ "Olympus-Athena/cmd/docs"
	"Olympus-Athena/pkg/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	swagger "github.com/swaggo/fiber-swagger"
	"log/slog"
)

// @title Athena API
// @version 1.0
// @description This service manages Identity and Access Management
func main() {
	fiberApp := fiber.New(fiber.Config{
		AppName: "Athena API",
	})

	// middleware
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
		athena.Get("/health", handlers.HealthHandler)
	}

	slog.Error("Critical Error launching API", "Error", fiberApp.Listen(":4200"))
}
