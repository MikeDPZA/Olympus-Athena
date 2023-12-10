package errors

import "Olympus-Athena/pkg/errors/codes"

type GrantTypeError struct {
	GrantType string
}

func (g GrantTypeError) Code() string {
	return codes.GrantTypeErrorCode
}

func (g GrantTypeError) Error() string {
	return "Invalid grant type given: " + g.GrantType
}
