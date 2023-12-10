package errors

import "Olympus-Athena/pkg/errors/codes"

type NoRedirectsGivenForGrantTypeError struct {
	GrantType string
}

func (n NoRedirectsGivenForGrantTypeError) Code() string {
	return codes.NoRedirectsGivenForGrantTypeErrorCode
}

func (n NoRedirectsGivenForGrantTypeError) Error() string {
	return "No redirect_uris given for grant_type that requires it: " + n.GrantType
}
