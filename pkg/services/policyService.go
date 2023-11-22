package services

import (
	"Olympus-Athena/pkg/responses"
	"github.com/google/uuid"
)

type PolicyService interface {
	GetPolicyById(id uuid.UUID) (responses.PolicyResponse, error)
}
