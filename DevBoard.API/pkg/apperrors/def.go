package apperrors

type ErrorDef struct {
	Code    string
	Message string
}

var (
	ErrUnknown          = ErrorDef{Code: "UNKNOWN_ERROR", Message: "An unknown error occurred"}
	ErrInvalidRequest   = ErrorDef{Code: "INVALID_REQUEST", Message: "Invalid request body"}
	ErrUnauthorized     = ErrorDef{Code: "UNAUTHORIZED", Message: "Unauthorized access"}
	ErrForbidden        = ErrorDef{Code: "FORBIDDEN", Message: "Access forbidden"}
	ErrNotFound         = ErrorDef{Code: "NOT_FOUND", Message: "Resource not found"}
	ErrInternalServer   = ErrorDef{Code: "INTERNAL_SERVER_ERROR", Message: "Internal server error"}
	ErrValidationFailed = ErrorDef{Code: "VALIDATION_FAILED", Message: "Validation failed"}

	// Auth
	ErrInvalidCredentials = ErrorDef{Code: "INVALID_CREDENTIALS", Message: "Invalid email or password"}
	ErrInvalidToken       = ErrorDef{Code: "INVALID_TOKEN", Message: "Invalid or expired token"}

	// User
	ErrUserNotFound      = ErrorDef{Code: "USER_NOT_FOUND", Message: "User not found"}
	ErrUserAlreadyExists = ErrorDef{Code: "USER_ALREADY_EXISTS", Message: "User already exists"}
)
