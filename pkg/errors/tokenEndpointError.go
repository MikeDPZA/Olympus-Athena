package errors

type TokenEndpointAuthMethodError struct {
	Method string
}

func (t TokenEndpointAuthMethodError) Error() string {
	return tokenEndpointAuthErrorCode + " - Invalid Token Endpoint Authentication method given: " + t.Method
}
