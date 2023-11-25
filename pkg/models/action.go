package models

import "github.com/google/uuid"

type Action struct {
	Base
	Name          string `gorm:"size:32"`
	Key           string `gorm:"size:32;unique"`
	FeatureId     uuid.UUID
	PolicyActions []PolicyAction `gorm:"foreignKey:ActionId"`
}
