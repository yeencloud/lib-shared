package apperr

import (
	"errors"

	"github.com/samber/lo"
)

type ErrorType string

const (
	ErrorTypeNotImplemented   ErrorType = "not_implemented"
	ErrorTypeInvalidArgument  ErrorType = "invalid_argument"
	ErrorTypeInternal         ErrorType = "internal_issue"
	ErrorTypeUnauthorized     ErrorType = "unauthorized"
	ErrorTypeConflict         ErrorType = "conflict"
	ErrorTypeUnavailable      ErrorType = "service_unavailable"
	ErrorTypeResourceNotFound ErrorType = "resource_not_found"
)

type TypedError interface {
	Type() ErrorType
}

func GetErrorTypeOrNil(err error) *ErrorType {
	var restError TypedError
	if err != nil && errors.As(err, &restError) {
		return lo.ToPtr(restError.Type())
	}

	return nil
}
