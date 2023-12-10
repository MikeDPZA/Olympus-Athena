package errors

import "Olympus-Athena/pkg/errors/codes"

type DoubleJwksError struct {
	errorBase
}

func (d DoubleJwksError) Code() string {
	return codes.DoubleJwksErrorCode
}

func (d DoubleJwksError) Error() string {
	return "Both Jwks and JwksUri was given, only provide one."
}
