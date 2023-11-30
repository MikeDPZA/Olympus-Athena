package validators

import (
	"Olympus-Athena/pkg/errors"
	"Olympus-Athena/pkg/requests/oauth"
)

func ValidateGrantTypes(gt []string) error {
	validGrants := map[string]bool{
		oauth.AuthorizationCodeGrantType: true,
		oauth.ImplicitGrantType:          true,
		oauth.PasswordGrantType:          true,
		oauth.ClientCredentialsGrantType: true,
		oauth.RefreshTokenGrantType:      true,
	}

	for _, grantType := range gt {
		if !validGrants[grantType] {
			return errors.GrantTypeError{GrantType: grantType}
		}
	}

	return nil
}
