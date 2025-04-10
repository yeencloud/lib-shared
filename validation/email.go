package validation

import (
	"context"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func emailValidator(ctx context.Context, fl validator.FieldLevel) error {
	value := fl.Field().String()

	if len(value) == 0 {
		return errors.New("email cannot be empty")
	}

	re := regexp.MustCompile(`(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`) // TODO: MustCompile shold be used only once for performance
	if !re.MatchString(value) {
		return errors.New("invalid email format")
	}

	return nil
}
