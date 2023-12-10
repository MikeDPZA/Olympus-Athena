package routes

import (
	"Olympus-Athena/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupOAuthRoutes(athena fiber.Router, clientController *controllers.OAuthClientController) {
	oauth := athena.Group("/oauth")
	{
		client := oauth.Group("/clients")
		{
			client.Put("/", clientController.PutClient)
		}
	}
}
