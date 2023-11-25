package models

import "github.com/google/uuid"

type ContainerPolicy struct {
	Base
	PolicyId    uuid.UUID
	ContainerId uuid.UUID
}
