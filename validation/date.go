package validation

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
)

func rfc3339Validator(ctx context.Context, fl validator.FieldLevel) error {
	dateStr := fl.Field().String()
	_, err := time.Parse(time.RFC3339, dateStr)
	return err
}
