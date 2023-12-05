package validators

import (
	"Olympus-Athena/pkg/errors"
	"Olympus-Athena/pkg/requests/oauth"
	"net/url"
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

func ValidateResponseType(rt string) error {
	validResponseTypes := map[string]bool{
		oauth.CodeResponseType:  true,
		oauth.TokenResponseType: true,
	}

	if !validResponseTypes[rt] {
		return errors.ResponseTypeError{ResponseType: rt}
	}

	return nil
}

func ValidateTokenEndpointAuth(m string) error {
	validTypes := map[string]bool{
		oauth.NoneTokenEndpointAuthMethod:              true,
		oauth.ClientSecretPostTokenEndpointAuthMethod:  true,
		oauth.ClientSecretBasicTokenEndpointAuthMethod: true,
	}

	if !validTypes[m] {
		return errors.TokenEndpointAuthMethodError{Method: m}
	}

	return nil
}

func ValidateRedirectUris(urls []string) error {
	for _, raw := range urls {
		_, err := url.Parse(raw)
		if err != nil {

		}
	}
	return nil
}
