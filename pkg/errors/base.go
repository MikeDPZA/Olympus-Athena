package errors

type errorBase struct {
	ErrorCode string
}

type AthenaError interface {
	error
	Code() string
}
