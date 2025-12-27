package domain

import "github.com/yeencloud/lib-shared/apperr"

type ValidationError struct {
	Source           error
	ValidationIssues map[string]string
}

func (e ValidationError) Error() string {
	return "validation failed"
}

func (e ValidationError) Unwrap() error {
	return apperr.InvalidArgumentError{}
}

func (e ValidationError) Details() apperr.ErrorDetails {
	return apperr.ErrorDetails{
		Reason:  e.Source.Error(),
		Details: e.ValidationIssues,
	}
}
