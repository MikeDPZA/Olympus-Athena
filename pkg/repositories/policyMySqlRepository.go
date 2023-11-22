package repositories

import (
	"Olympus-Athena/pkg/database"
	"Olympus-Athena/pkg/models"
	"github.com/google/uuid"
	"log/slog"
)

type PolicyMySqlRepository struct {
	db *database.AthenaDb
}

func NewPolicyMySqlRepository(db *database.AthenaDb) *PolicyMySqlRepository {
	return &PolicyMySqlRepository{
		db: db,
	}
}

func (p *PolicyMySqlRepository) FindById(id uuid.UUID) (*models.Policy, error) {
	slog.Info("Finding policy")
	return nil, nil
}
