package models

import "github.com/google/uuid"

type UserPolicy struct {
	Base
	UserId   uuid.UUID
	PolicyId uuid.UUID
}
