package models

import "gorm.io/gorm"

type Policy struct {
	gorm.Model
	Name string
}
