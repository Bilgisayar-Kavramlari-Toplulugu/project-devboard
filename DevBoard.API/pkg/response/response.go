package response

import (
	"net/http"

	"project-devboard/pkg/apperrors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *ErrorBody  `json:"error"`
}

type ErrorBody struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type MessageBody struct {
	Message string `json:"message"`
}

func Success(c *gin.Context, statusCode int, data interface{}) {
	if statusCode == http.StatusNoContent {
		c.Status(statusCode)
		return
	}

	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func Message(c *gin.Context, statusCode int, message string) {
	Success(c, statusCode, MessageBody{Message: message})
}

func Error(c *gin.Context, statusCode int, code, message string, details interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Data:    nil,
		Error: &ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

func ErrorWithCode(c *gin.Context, statusCode int, def apperrors.ErrorDef) {
	Error(c, statusCode, def.Code, def.Message, nil)
}
