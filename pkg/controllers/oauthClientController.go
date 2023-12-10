package controllers

import (
	"Olympus-Athena/pkg/requests/oauth"
	"Olympus-Athena/pkg/responses"
	"Olympus-Athena/pkg/sanitizers"
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

// PutClient godoc
// @Summary Create an OAuth Client
// @Tags OAuth Client
// @Param request body oauth.PutOAuthClientRequest true "Body to create a client"
// @Success 200 {object} responses.PutOAuthClientResponse
// @Failure 400 {object} responses.ApiError
// @Router /athena/v1/oauth/clients [put]
func (o OAuthClientController) PutClient(ctx *fiber.Ctx) error {
	var request oauth.PutOAuthClientRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		slog.Error("Couldn't parse body", "error", err)
		return BadRequest(ctx, responses.ParseToApiError(err))
	}

	sanitizers.SanitizePutOAuthClientRequest(&request)

	vErr := validators.ValidatePutOAuthClientRequest(&request)
	if vErr != nil {
		return BadRequest(ctx, responses.NewApiError(vErr))
	}

	randomBytes := make([]byte, 128)
	_, err = rand.Read(randomBytes)
	if err != nil {
		slog.Error("Couldn't create secret", "Error", err)
		return BadRequest(ctx, responses.ParseToApiError(err))
	}

	clientId := uuid.New()
	secret := base64.URLEncoding.EncodeToString(randomBytes)

	return Ok(ctx, responses.PutOAuthClientResponse{
		ClientId:     clientId,
		ClientSecret: secret,
	})
}
