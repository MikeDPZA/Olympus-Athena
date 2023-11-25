package routes

import (
	"Olympus-Athena/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupPolicyRoutes(athenaGroup fiber.Router, policyHandler *controllers.PolicyController) {
	policy := athenaGroup.Group("/policies")
	{
		policy.Get("/:id", policyHandler.GetPolicyById)
	}
}
