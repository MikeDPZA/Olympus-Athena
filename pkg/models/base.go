package models

import (
	"github.com/google/uuid"
	"time"
)

type Base struct {
	Id        uuid.UUID `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time `gorm:"autoUpdateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}
