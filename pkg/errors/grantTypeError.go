package errors

type GrantTypeError struct {
	GrantType string
}

func (g GrantTypeError) Error() string {
	return "Invalid grant type given: " + g.GrantType
}
