package repositories

import (
	"Olympus-Athena/pkg/models"
	"github.com/google/uuid"
)

type PolicyRepository interface {
	FindById(id uuid.UUID) (*models.Policy, error)
}
