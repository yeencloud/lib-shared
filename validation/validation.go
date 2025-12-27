package validation

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/yeencloud/lib-shared/domain"
)

type Validator struct {
	*validator.Validate
}

type ValidationFunc func(ctx context.Context, fl validator.FieldLevel) error

func (v *Validator) RegisterValidationFunc(name string, fn ValidationFunc) error {
	return v.RegisterValidationCtx(name, func(ctx context.Context, fl validator.FieldLevel) bool {
		err := fn(ctx, fl)
		return err == nil
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

func (v *Validator) ValidateStruct(obj interface{}) error {
	err := v.Validate.Struct(obj)
	if err == nil {
		return nil
	}

	return v.validateReq(obj)
}

func (v *Validator) validateReq(req any) error {
	if err := v.Validate.Struct(req); err != nil {
		var validationError *validator.InvalidValidationError
		if errors.As(err, &validationError) {
			return err
		}

		var validationErrors validator.ValidationErrors
		if !errors.As(err, &validationErrors) {
			return err
		}

		maps := map[string]string{}

		for _, fe := range validationErrors {
			tag := fe.Tag()
			value := fe.Param()
			if value != "" {
				tag = tag + "=" + value
			}
			maps[fe.Field()] = tag
		}

		return domain.ValidationError{
			Source:           err,
			ValidationIssues: maps,
		}
	}
	return nil
}

func NewValidator() (*Validator, error) {
	v := validator.New()

	validatorEngine := &Validator{v}

	validations := map[string]ValidationFunc{
		"date_time":       rfc3339Validator,
		"uuid":            uuidValidator,
		"startswithalpha": startswithalphaValidator,
	}

	err := validatorEngine.RegisterValidations(validations)

	if err != nil {
		return nil, err
	}

	return validatorEngine, nil
}
