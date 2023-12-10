package models

type OAuthClient struct {
	Base
	Name                    string `gorm:"size:32"`
	ClientUri               string `gorm:"size:512"`
	Contacts                string
	GrantTypes              string `gorm:"size:128"`
	LogoUri                 string `gorm:"size:512"`
	Jwks                    string `gorm:"size:2048"`
	JwksUri                 string `gorm:"size:512"`
	RedirectUris            string `gorm:"size:2048"`
	PolicyUri               string `gorm:"size:512"`
	ResponseType            string `gorm:"size:16"`
	Scope                   string `gorm:"size:512"`
	SoftwareId              string `gorm:"size:128"`
	SoftwareVersion         string `gorm:"size:64"`
	TokenEndpointAuthMethod string `gorm:"size:32"`
	TosUri                  string `gorm:"size:512"`
}
