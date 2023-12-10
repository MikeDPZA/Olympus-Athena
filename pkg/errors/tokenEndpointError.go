package errors

import "Olympus-Athena/pkg/errors/codes"

type TokenEndpointAuthMethodError struct {
	Method string
}

func (t TokenEndpointAuthMethodError) Code() string {
	return codes.TokenEndpointAuthErrorCode
}

func (t TokenEndpointAuthMethodError) Error() string {
	return "Invalid Token Endpoint Authentication method given: " + t.Method
}
