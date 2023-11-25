package models

import "github.com/google/uuid"

type PolicyAction struct {
	Base
	ActionId uuid.UUID
	PolicyId uuid.UUID
}
