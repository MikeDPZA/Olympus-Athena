package validators

import (
	"Olympus-Athena/pkg/requests/oauth"
	"log/slog"
)

func ValidateClientRegistrationRequest(r *oauth.ClientRegistrationRequest) error {
	if err := ValidateGrantTypes(r.GrantTypes); err != nil {
		slog.Error("Couldn't validate grant_type", "error", err)
		return err
	}

	if err := ValidateResponseType(r.ResponseType); err != nil {
		slog.Error("Couldn't validate response_type", "error", err)
		return err
	}

	if err := ValidateTokenEndpointAuth(r.TokenEndpointAuthMethod); err != nil {
		slog.Error("Couldn't validate token_endpoint_auth_method", "error", err)
		return err
	}

	if err := ValidateRedirectUris(r.RedirectUris); err != nil {
		slog.Error("Couldn't validate redirect_uris", "error", err)
		return err
	}

	return nil
}
