package validation

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func uuidValidator(ctx context.Context, fl validator.FieldLevel) error {
	value := fl.Field().String()
	_, err := uuid.Parse(value)
	return err
}
