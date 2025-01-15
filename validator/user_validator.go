package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		validator *validator.Validate
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}
)

var Valid = validator.New()

var MyValidator = &XValidator{
	validator: Valid,
}

func (v XValidator) Validate(data interface{}) []ErrorResponse {
	validationError := []ErrorResponse{}

	errs := Valid.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Error = true
			elem.Value = err.Value()

			validationError = append(validationError, elem)
		}
	}

	return validationError
}

func init() {
	MyValidator.validator.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		pattern := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		return pattern.MatchString(fl.Field().String())
	})
}
