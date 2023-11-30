package controllers

import (
	"Olympus-Athena/pkg/requests/oauth"
	"Olympus-Athena/pkg/responses"
	"Olympus-Athena/pkg/validators"
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log/slog"
)

type OAuthClientController struct {
}

func NewOAuthClientController() *OAuthClientController {
	return &OAuthClientController{}
}

// RegisterClient godoc
// @Summary Create an OAuth Client
// @Tags OAuth Client
// @Param request body oauth.ClientRegistrationRequest true "Body to create a client"
// @Router /athena/v1/oauth/clients [post]
func (o OAuthClientController) RegisterClient(ctx *fiber.Ctx) error {
	var request oauth.ClientRegistrationRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		slog.Error("Couldn't parse body", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(&responses.ApiError{
			Error: err.Error(),
		})
	}

	err = validators.ValidateGrantTypes(request.GrantTypes)
	if err != nil {
		slog.Error("Couldn't validate grant_type", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(&responses.ApiError{
			Error: err.Error(),
		})
	}

	randomBytes := make([]byte, 128)
	_, err = rand.Read(randomBytes)
	if err != nil {
		slog.Error("Couldn't create secret", "Error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(&responses.ApiError{
			Error: err.Error(),
		})
	}

	clientId := uuid.New()
	secret := base64.URLEncoding.EncodeToString(randomBytes)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"client_id": clientId,
		"secret":    secret,
	})
}
