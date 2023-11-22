package services

import (
	"Olympus-Athena/pkg/repositories"
	"Olympus-Athena/pkg/responses"
	"github.com/google/uuid"
)

type PolicyMySqlService struct {
	policyRepository repositories.PolicyRepository
}

func NewPolicyMySqlService(policyRepository repositories.PolicyRepository) *PolicyMySqlService {
	return &PolicyMySqlService{
		policyRepository: policyRepository,
	}
}

func (p PolicyMySqlService) GetPolicyById(id uuid.UUID) (responses.PolicyResponse, error) {
	//policy, err := p.policyRepository.FindById(id)
	panic("implement me")
}
