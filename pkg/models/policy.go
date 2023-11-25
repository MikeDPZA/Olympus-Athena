package models

type Policy struct {
	Base
	Name              string            `gorm:"size:32"`
	PolicyActions     []PolicyAction    `gorm:"foreignKey:PolicyId"`
	ContainerPolicies []ContainerPolicy `gorm:"foreignKey:PolicyId"`
	UserPolicies      []UserPolicy      `gorm:"foreignKey:PolicyId"`
}
