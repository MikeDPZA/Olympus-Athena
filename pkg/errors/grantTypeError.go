package errors

type GrantTypeError struct {
	GrantType string
}

func (g GrantTypeError) Error() string {
	return grantTypeErrorCode + " - Invalid grant type given: " + g.GrantType
}
