package errors

type InvalidUrlGivenError struct {
	Url string
}

func (i InvalidUrlGivenError) Error() string {
	return invalidUrlGivenErrorCode + " - Invalid url was provided: " + i.Url
}
