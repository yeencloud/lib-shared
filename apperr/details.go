package apperr

type ErrorDetails struct {
	Reason  string
	Details map[string]string
}

type DetailedError interface {
	Details() ErrorDetails
}
