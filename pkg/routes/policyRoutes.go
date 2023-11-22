package routes

import (
	"Olympus-Athena/pkg/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupPolicyRoutes(athenaGroup fiber.Router, policyHandler *handlers.PolicyHandler) {
	policy := athenaGroup.Group("/policies")
	{
		policy.Get("/:id", policyHandler.HandleGetPolicyById)
	}
}
