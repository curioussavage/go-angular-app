package validation

import (
	"fmt"
	"reflect"

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
		for idx := range validationErrors {
			err := validationErrors[idx]
			field := getJSONTag(i, err.Field())
			message := fmt.Sprintf("Validation failed for %s", field)
			appErrors = append(appErrors, apperrors.ValidationError{Field: field, Message: message})
		}
		// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return appErrors
	}
	return nil
}

func getJSONTag(structObj interface{}, fieldName string) string {
	typ := reflect.TypeOf(structObj)
	field, found := typ.FieldByName(fieldName)
	if !found {
		return ""
	}
	return field.Tag.Get("json")
}
