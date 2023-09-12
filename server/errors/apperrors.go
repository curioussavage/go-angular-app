package apperrors

import "fmt"

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return "errors"
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}
