package apperrors

type ErrorCode string

const (
    NotFound       ErrorCode = "NOT_FOUND"
    BadRequest     ErrorCode = "BAD_REQUEST"
    Conflict       ErrorCode = "CONFLICT"
    Unauthorized   ErrorCode = "UNAUTHORIZED"
    Forbidden      ErrorCode = "FORBIDDEN"
    ValidationErr  ErrorCode = "VALIDATION_FAILED"
    InvalidRequest ErrorCode = "INVALID_REQUEST"
    InternalError  ErrorCode = "INTERNAL_SERVER_ERROR"
    InvalidCreds   ErrorCode = "INVALID_CREDENTIALS"
    InvalidToken   ErrorCode = "INVALID_TOKEN"
)