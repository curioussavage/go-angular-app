package validation

import (
	"fmt"

	apperrors "github.com/curioussavage/integra/errors"
	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var appErrors apperrors.ValidationErrors
		for i := range validationErrors {
			err := validationErrors[i]
			message := fmt.Sprintf("Validation failed for %s", err.Tag())
			appErrors = append(appErrors, apperrors.ValidationError{Field: err.Tag(), Message: message})
		}
		// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return appErrors
	}
	return nil
}
