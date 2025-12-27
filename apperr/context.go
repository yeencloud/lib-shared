package apperr

type ObjectNotInContextError struct {
	Object string
}

func (e *ObjectNotInContextError) Error() string {
	return e.Object + " not in context"
}

type WrongObjectTypeInContextError struct {
	Object       string
	ExpectedType string
}

func (e *WrongObjectTypeInContextError) Error() string {
	return "wrong object type in context: " + e.Object + " should be " + e.ExpectedType
}
