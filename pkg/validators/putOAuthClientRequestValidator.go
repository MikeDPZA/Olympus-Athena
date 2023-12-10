package validators

import (
	"Olympus-Athena/pkg/errors"
	"Olympus-Athena/pkg/requests/oauth"
	"log/slog"
)

func ValidatePutOAuthClientRequest(r *oauth.PutOAuthClientRequest) errors.AthenaError {
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

	if err := ValidateGrantTypesHaveRedirectUri(r.GrantTypes, r.RedirectUris); err != nil {
		slog.Error("No redirect_uris provided for given grant_types", "error", err)
		return err
	}

	if err := ValidateJwks(r); err != nil {
		slog.Error("Couldn't validate jwks and jwks_uri", "error", err)
		return err
	}

	if err := ValidateConflictingGrantTypes(r.GrantTypes); err != nil {
		slog.Error("Conflicting grant_types", "error", err)
		return err
	}

	return nil
}
