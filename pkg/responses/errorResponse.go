package responses

import (
	"Olympus-Athena/pkg/errors"
	"Olympus-Athena/pkg/errors/codes"
)

type ApiError struct {
	ErrorCode string `json:"error_code"`
	Error     string `json:"error"`
}

func NewApiError(e errors.AthenaError) *ApiError {
	return &ApiError{
		Error:     e.Error(),
		ErrorCode: e.Code(),
	}
}

func ParseToApiError(e error) *ApiError {
	return &ApiError{
		ErrorCode: codes.DefaultErrorCode,
		Error:     e.Error(),
	}
}
