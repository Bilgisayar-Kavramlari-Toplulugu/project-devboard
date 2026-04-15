package handler

import (
	"net/http"

	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service   services.UserService
	validator *validator.Validator
}

func NewUserHandler(service services.UserService, validator *validator.Validator) *UserHandler {
	return &UserHandler{service: service, validator: validator}
}

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.UserCreateRequest  true  "Create User Request"
// @Success      201      {object}  UserEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Unauthorized"
// @Failure      409      {object}  response.Response  "User already exists"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req dtos.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}

	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}

	userID := userIDFromContext(c, uuid.Nil)

	user, err := h.service.CreateUser(c.Request.Context(), req, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusCreated, dtos.NewUserResponse(user))
}

// GetByID godoc
// @Summary      Get a user by ID
// @Description  Get details of a specific user by UUID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User UUID"
// @Success      200  {object}  UserEnvelope
// @Failure      400  {object}  response.Response  "Invalid ID"
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      404  {object}  response.Response  "User not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.BadRequest, apperrors.ErrInvalidRequest))
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusOK, dtos.NewUserResponse(user))
}

// List godoc
// @Summary      List users
// @Description  Get a list of users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Limit (default 10)"
// @Param        offset  query     int  false  "Offset (default 0)"
// @Success      200     {object}  UserListEnvelope
// @Failure      401     {object}  response.Response  "Unauthorized"
// @Failure      500     {object}  response.Response  "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) List(c *gin.Context) {
	limit, offset := paginationParams(c)

	users, err := h.service.ListUsers(limit, offset)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusOK, dtos.NewUserResponses(users))
}

// Update godoc
// @Summary      Update a user
// @Description  Update a user's details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id       path      string                 true  "User UUID"
// @Param        request  body      dtos.UserUpdateRequest true  "Update User Request"
// @Success      200      {object}  UserEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Unauthorized"
// @Failure      404      {object}  response.Response  "User not found"
// @Failure      409      {object}  response.Response  "User already exists"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.BadRequest, apperrors.ErrInvalidRequest))
		return
	}

	var req dtos.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}

	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}

	userID := userIDFromContext(c, id)

	user, err := h.service.UpdateUser(c.Request.Context(), id, req, userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusOK, dtos.NewUserResponse(user))
}

// Delete godoc
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User UUID"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  response.Response  "Invalid ID"
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      404  {object}  response.Response  "User not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.BadRequest, apperrors.ErrInvalidRequest))
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusNoContent, nil)
}

// helpers
func toAppFieldErrors(errs []validator.ValidationError) []apperrors.FieldError {
	result := make([]apperrors.FieldError, len(errs))
	for i, e := range errs {
		result[i] = apperrors.FieldError{Field: e.Field, Message: e.Message}
	}
	return result
}
