package controllers

import (
	"Olympus-Athena/pkg/responses"
	"Olympus-Athena/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PolicyController struct {
	policyService services.PolicyService
}

func NewPolicyController(ps services.PolicyService) *PolicyController {
	return &PolicyController{
		policyService: ps,
	}
}

// GetPolicyById godoc
// @Summary Get Policy by ID
// @Tags Policies
// @Param id path string true "ID of Policy"
// @Router /athena/v1/policies/{id} [get]
func (ph *PolicyController) GetPolicyById(ctx *fiber.Ctx) error {
	idString := ctx.Params("id")

	id := uuid.MustParse(idString)

	policy, err := ph.policyService.GetPolicyById(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(responses.ApiError{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(policy)
}
