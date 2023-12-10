package errors

import "Olympus-Athena/pkg/errors/codes"

type InvalidUrlGivenError struct {
	Url string
}

func (i InvalidUrlGivenError) Code() string {
	return codes.InvalidUrlGivenErrorCode
}

func (i InvalidUrlGivenError) Error() string {
	return "Invalid url was provided: " + i.Url
}
