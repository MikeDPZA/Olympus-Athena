package errors

type ResponseTypeError struct {
	ResponseType string
}

func (r ResponseTypeError) Error() string {
	return responseTypeErrorCode + " - Invalid response type given: " + r.ResponseType
}
