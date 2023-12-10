package sanitizers

import "Olympus-Athena/pkg/requests/oauth"

func SanitizePutOAuthClientRequest(req *oauth.PutOAuthClientRequest) {
	if len(req.GrantTypes) == 0 {
		req.GrantTypes = append(req.GrantTypes, oauth.AuthorizationCodeGrantType)
	}

	if len(req.ResponseType) == 0 {
		req.ResponseType = oauth.CodeResponseType
	}
}
