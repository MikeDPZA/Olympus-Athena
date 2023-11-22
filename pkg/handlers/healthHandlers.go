package handlers

import (
	"Olympus-Athena/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"time"
)

// HealthHandler godoc
// @Summary Component Health Check
// @Description Check if component is up and running
// @Success 200 {object} responses.HealthResponse
// @Tags Health
// @Router /athena/health [get]
func HealthHandler(ctx *fiber.Ctx) error {
	slog.Info("Athena is healthy")
	return ctx.Status(fiber.StatusOK).JSON(responses.HealthResponse{
		Success: true,
		Time:    time.Now(),
	})
}
