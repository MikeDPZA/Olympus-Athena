package errors

import "Olympus-Athena/pkg/errors/codes"

type ResponseTypeError struct {
	ResponseType string
}

func (r ResponseTypeError) Code() string {
	return codes.ResponseTypeErrorCode
}

func (r ResponseTypeError) Error() string {
	return "Invalid response type given: " + r.ResponseType
}
