package models

type Container struct {
	Base
	Name              string            `gorm:"size:32"`
	ContainerPolicies []ContainerPolicy `gorm:"foreignKey:ContainerId"`
	UserContainers    []UserContainer   `gorm:"foreignKey:ContainerId"`
}
