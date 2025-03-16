package errors

type FixableError interface {
	HowToFix() string
}
