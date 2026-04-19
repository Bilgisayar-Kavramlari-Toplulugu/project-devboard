package middleware

import (
	"errors"
	"log"
	"net/http"
	"runtime/debug"

	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"

	"github.com/gin-gonic/gin"
)

var errorStatusMap = map[apperrors.ErrorCode]int{
	apperrors.NotFound:       http.StatusNotFound,
	apperrors.BadRequest:     http.StatusBadRequest,
	apperrors.ValidationErr:  http.StatusBadRequest,
	apperrors.Conflict:       http.StatusConflict,
	apperrors.Unauthorized:   http.StatusUnauthorized,
	apperrors.Forbidden:      http.StatusForbidden,
	apperrors.InvalidRequest: http.StatusBadRequest,
	apperrors.InvalidCreds:   http.StatusUnauthorized,
	apperrors.InvalidToken:   http.StatusUnauthorized,
	apperrors.InternalError:  http.StatusInternalServerError,
}

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Printf("panic recovered: %v\n%s", recovered, debug.Stack())
				if !c.Writer.Written() {
					response.ErrorWithCode(c, http.StatusInternalServerError, apperrors.ErrInternalServer)
				}
				c.Abort()
			}
		}()

		c.Next()

		if len(c.Errors) == 0 || c.Writer.Written() {
			return
		}

		err := c.Errors.Last().Err

		var appErr *apperrors.AppError
		if errors.As(err, &appErr) {
			statusCode, ok := errorStatusMap[appErr.Kind]
			if !ok {
				statusCode = http.StatusInternalServerError
			}

			var details interface{}
			if len(appErr.FieldErrors) > 0 {
				details = appErr.FieldErrors
			}

			response.Error(c, statusCode, appErr.Code, appErr.Message, details)
			return
		}

		response.ErrorWithCode(c, http.StatusInternalServerError, apperrors.ErrInternalServer)
	}
}
