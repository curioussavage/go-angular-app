package apperrors

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	return "errors"
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
