package errors

type IdentifiableError interface {
	Identifier() string
}
