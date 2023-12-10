package controllers

import (
	"Olympus-Athena/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

func BadRequest(ctx *fiber.Ctx, r *responses.ApiError) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(&r)
}

func Ok(ctx *fiber.Ctx, r any) error {
	return ctx.Status(fiber.StatusOK).JSON(&r)
}
