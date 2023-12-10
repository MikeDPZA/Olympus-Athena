package errors

import (
	"Olympus-Athena/pkg/errors/codes"
	"Olympus-Athena/pkg/requests/oauth"
)

type ConflictingGrantTypesError struct {
}

func (c ConflictingGrantTypesError) Error() string {
	return codes.ConflictingGrantTypesErrorCode
}

func (c ConflictingGrantTypesError) Code() string {
	return "Conflicting grant_types were given: " + oauth.AuthorizationCodeGrantType + ", " + oauth.ImplicitGrantType
}
