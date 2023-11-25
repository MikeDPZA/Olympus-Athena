package models

type User struct {
	Base
	FirstName      string `gorm:"size:32"`
	LastName       string `gorm:"size:32"`
	Identifier     string `gorm:"unique;size:64"`
	Password       string
	UserPolicies   []UserPolicy    `gorm:"foreignKey:UserId"`
	UserContainers []UserContainer `gorm:"foreignKey:UserId"`
}
