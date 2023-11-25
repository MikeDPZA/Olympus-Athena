package models

type Feature struct {
	Base
	Name    string   `gorm:"size:32"`
	Actions []Action `gorm:"foreignKey:FeatureId"`
}
