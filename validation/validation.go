package validation

import (
	"context"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

type Validator struct {
	*validator.Validate
}

type ValidationFunc func(ctx context.Context, fl validator.FieldLevel) error

func (v *Validator) RegisterValidationFunc(name string, fn ValidationFunc) error {
	return v.RegisterValidationCtx(name, func(ctx context.Context, fl validator.FieldLevel) bool {
		err := fn(ctx, fl)
		if err != nil {
			log.WithContext(ctx).WithField("tag", name).WithError(err).Warn("Validation failed")
			return false
		}
		return true
	})
}

func (v *Validator) RegisterValidations(validations map[string]ValidationFunc) error {
	for name, fn := range validations {
		if err := v.RegisterValidationFunc(name, fn); err != nil {
			return err
		}
	}
	return nil
}

func NewValidator() (*Validator, error) {
	v := validator.New()

	validatorEngine := &Validator{v}

	validations := map[string]ValidationFunc{
		"date_time": rfc3339Validator,
		"uuid":      uuidValidator,
		"email":     emailValidator,
	}

	for name, fn := range validations {
		if err := validatorEngine.RegisterValidationFunc(name, fn); err != nil {
			return nil, err
		}
	}

	return validatorEngine, nil
}
