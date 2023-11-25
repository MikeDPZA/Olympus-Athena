package models

import "github.com/google/uuid"

type UserContainer struct {
	Base
	UserId      uuid.UUID
	ContainerId uuid.UUID
}
