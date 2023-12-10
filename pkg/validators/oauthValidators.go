package validators

import (
	"Olympus-Athena/pkg/errors"
	"Olympus-Athena/pkg/requests/oauth"
	"net/url"
	"slices"
)

func ValidateGrantTypes(gt []string) errors.AthenaError {
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

func ValidateResponseType(rt string) errors.AthenaError {
	validResponseTypes := map[string]bool{
		oauth.CodeResponseType:  true,
		oauth.TokenResponseType: true,
	}

	if !validResponseTypes[rt] {
		return errors.ResponseTypeError{ResponseType: rt}
	}

	return nil
}

func ValidateTokenEndpointAuth(m string) errors.AthenaError {
	validTypes := map[string]bool{
		oauth.NoneTokenEndpointAuthMethod:              true,
		oauth.ClientSecretPostTokenEndpointAuthMethod:  true,
		oauth.ClientSecretBasicTokenEndpointAuthMethod: true,
	}

	if len(m) > 0 && !validTypes[m] {
		return errors.TokenEndpointAuthMethodError{Method: m}
	}

	return nil
}

func ValidateRedirectUris(urls []string) errors.AthenaError {
	for _, raw := range urls {
		_, err := url.Parse(raw)
		if err != nil {
			return errors.InvalidUrlGivenError{Url: raw}
		}
	}
	return nil
}

func ValidateJwks(req *oauth.PutOAuthClientRequest) errors.AthenaError {
	if req.JwksUri != "" && req.Jwks != nil {
		return errors.DoubleJwksError{}
	}

	return nil
}

func ValidateGrantTypesHaveRedirectUri(gts []string, uris []string) errors.AthenaError {
	grantTypesRequireRedirects := map[string]bool{
		oauth.AuthorizationCodeGrantType: true,
		oauth.ImplicitGrantType:          true,
	}

	urisIsEmpty := len(uris) == 0

	for _, gt := range gts {
		if grantTypesRequireRedirects[gt] && urisIsEmpty {
			return errors.NoRedirectsGivenForGrantTypeError{GrantType: gt}
		}
	}

	return nil
}

func ValidateConflictingGrantTypes(gts []string) errors.AthenaError {
	if slices.Contains(gts, oauth.AuthorizationCodeGrantType) && slices.Contains(gts, oauth.ImplicitGrantType) {
		return errors.ConflictingGrantTypesError{}
	}

	return nil
}
