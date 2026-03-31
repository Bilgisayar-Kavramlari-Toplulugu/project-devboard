package handler

import (
	"net/http"
	"strconv"

	"project-devboard/internal/domain/entities"
	"project-devboard/internal/services"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service   services.UserService
	validator *validator.Validator
}

func NewUserHandler(service services.UserService, validator *validator.Validator) *UserHandler {
	return &UserHandler{
		service:   service,
		validator: validator,
	}
}

type CreateUserRequest struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required,min=6"`
	Firstname        string `json:"firstname" validate:"required,min=2,max=100"`
	Lastname         string `json:"lastname" validate:"required,min=2,max=100"`
	IsEmailValidated bool   `json:"is_email_validated" validate:"required"`
}

// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with the input payload
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body CreateUserRequest true "Create User Request"
// @Success      201  {object}  entities.User "Created User"
// @Failure      400  {object}  response.Response "Bad Request"
// @Failure      500  {object}  response.Response "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	userID := uuid.New()
	user := &entities.User{
		ID:               userID,
		Email:            req.Email,
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		Password:         string(hashedPassword), // Store hashed password
		IsEmailValidated: req.IsEmailValidated,
		BaseEntity: entities.BaseEntity{
			IsActive:       true,
			CreatedBy:      userID,
			LastModifiedBy: userID,
		},
	}

	if err := h.service.CreateUser(user); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Hide password in response
	user.Password = ""

	response.Success(c, http.StatusCreated, user)
}

// GetByID godoc
// @Summary      Get a user by ID
// @Description  Get details of a specific user by UUID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User UUID"
// @Success      200  {object}  entities.User
// @Failure      400  {object}  response.Response "Invalid ID"
// @Failure      404  {object}  response.Response "User not found"
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	// Hide password in response
	user.Password = ""

	response.Success(c, http.StatusOK, user)
}

// List godoc
// @Summary      List users
// @Description  Get a list of users with pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Limit (default 10)"
// @Param        offset  query     int  false  "Offset (default 0)"
// @Success      200  {array}   entities.User
// @Failure      500  {object}  response.Response "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	users, err := h.service.ListUsers(limit, offset)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	// Hide passwords in response
	for i := range users {
		users[i].Password = ""
	}

	response.Success(c, http.StatusOK, users)
}

// Update godoc
// @Summary      Update a user
// @Description  Update a user's details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id      path      string             true  "User UUID"
// @Param        request body      CreateUserRequest  true  "Update User Request"
// @Success      200     {object}  entities.User
// @Failure      400     {object}  response.Response "Bad Request"
// @Failure      500     {object}  response.Response "Internal Server Error"
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	ID, err := uuid.Parse(c.Param("ID"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Hash the password before updating if provided
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := &entities.User{
		ID:        ID,
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  string(hashedPassword), // Store hashed password
	}

	if err := h.service.UpdateUser(user); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Hide password in response
	user.Password = ""

	response.Success(c, http.StatusOK, user)
}

// Delete godoc
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User UUID"
// @Success      200  {object}  map[string]string "Success message"
// @Failure      400  {object}  response.Response "Invalid ID"
// @Failure      500  {object}  response.Response "Internal Server Error"
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
