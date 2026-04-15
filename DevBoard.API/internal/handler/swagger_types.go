package handler

import (
	"project-devboard/internal/dtos"
	"project-devboard/pkg/response"
)

// UserEnvelope documents the standard success envelope for a single user payload.
type UserEnvelope struct {
	Success bool                `json:"success"`
	Data    dtos.UserResponse   `json:"data"`
	Error   *response.ErrorBody `json:"error,omitempty"`
}

// UserListEnvelope documents the standard success envelope for user lists.
type UserListEnvelope struct {
	Success bool                `json:"success"`
	Data    []dtos.UserResponse `json:"data"`
	Error   *response.ErrorBody `json:"error,omitempty"`
}

// MessageEnvelope documents the standard success envelope for message responses.
type MessageEnvelope struct {
	Success bool                 `json:"success"`
	Data    response.MessageBody `json:"data"`
	Error   *response.ErrorBody  `json:"error,omitempty"`
}
