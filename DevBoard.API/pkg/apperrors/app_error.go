package apperrors

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type AppError struct {
	Kind        ErrorCode    `json:"-"`
	Code        string       `json:"code"`
	Message     string       `json:"message"`
	FieldErrors []FieldError `json:"errors,omitempty"`
	Err         error        `json:"-"`
}

func (e *AppError) Error() string { return e.Message }
func (e *AppError) Unwrap() error { return e.Err }

func New(kind ErrorCode, def ErrorDef) *AppError {
	return &AppError{Kind: kind, Code: def.Code, Message: def.Message}
}

func Wrap(kind ErrorCode, def ErrorDef, err error) *AppError {
	return &AppError{Kind: kind, Code: def.Code, Message: def.Message, Err: err}
}

func WithFieldErrors(kind ErrorCode, def ErrorDef, fieldErrors []FieldError) *AppError {
	return &AppError{
		Kind:        kind,
		Code:        def.Code,
		Message:     def.Message,
		FieldErrors: fieldErrors,
	}
}

func Validation(fieldErrors []FieldError) *AppError {
	return WithFieldErrors(ValidationErr, ErrValidationFailed, fieldErrors)
}
