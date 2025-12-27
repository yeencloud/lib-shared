package validation

import (
	"context"
	"errors"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func startswithalphaValidator(ctx context.Context, fl validator.FieldLevel) error {
	val := fl.Field().String()
	if len(val) > 0 && unicode.IsLetter(rune(val[0])) {
		return nil
	}
	return errors.New("must start with a letter")
}
