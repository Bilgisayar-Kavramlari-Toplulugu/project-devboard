package handler

import (
	"net/http"
	"strconv"

	"project-devboard/internal/domain/entities"
	"project-devboard/internal/services"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRoleHandler struct {
	service   services.UserRoleService
	validator *validator.Validator
}

func NewUserRoleHandler(service services.UserRoleService, v *validator.Validator) *UserRoleHandler {
	return &UserRoleHandler{
		service:   service,
		validator: v,
	}
}

// CreateUserRoleRequest - UserRole oluşturma request
type CreateUserRoleRequest struct {
	RoleID uuid.UUID `json:"role_id" validate:"required"`
}

// UpdateUserRoleRequest - UserRole güncelleme request
type UpdateUserRoleRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	RoleID uuid.UUID `json:"role_id" validate:"required"`
}

// @Summary Create a new user role assignment
// @Tags user-roles
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userRole body CreateUserRoleRequest true "UserRole data"
// @Success 201 {object} entities.UserRole
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /user-roles [post]
func (h *UserRoleHandler) CreateUserRole(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req CreateUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRole := &entities.UserRole{
		UserID: userID.(uuid.UUID),
		RoleID: req.RoleID,
	}

	if err := h.service.CreateUserRole(userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userRole)
}

// @Summary Get user role by ID
// @Tags user-roles
// @Produce json
// @Param id path string true "UserRole ID"
// @Success 200 {object} entities.UserRole
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user-roles/{id} [get]
func (h *UserRoleHandler) GetUserRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	userRole, err := h.service.GetUserRole(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user role not found"})
		return
	}

	c.JSON(http.StatusOK, userRole)
}

// @Summary List all user roles
// @Tags user-roles
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} entities.UserRole
// @Router /user-roles [get]
func (h *UserRoleHandler) ListUserRoles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	userRoles, err := h.service.ListUserRoles(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRoles)
}

// @Summary Update user role
// @Tags user-roles
// @Accept json
// @Produce json
// @Param id path string true "UserRole ID"
// @Param userRole body UpdateUserRoleRequest true "UserRole data"
// @Success 200 {object} entities.UserRole
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user-roles/{id} [put]
func (h *UserRoleHandler) UpdateUserRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var req UpdateUserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRole := &entities.UserRole{
		ID:     id,
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	if err := h.service.UpdateUserRole(userRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRole)
}

// @Summary Delete user role
// @Tags user-roles
// @Param id path string true "UserRole ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /user-roles/{id} [delete]
func (h *UserRoleHandler) DeleteUserRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	if err := h.service.DeleteUserRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get current user's roles
// @Tags user-roles
// @Produce json
// @Security BearerAuth
// @Success 200 {array} entities.UserRole
// @Failure 401 {object} map[string]interface{}
// @Router /user-roles/me [get]
func (h *UserRoleHandler) GetMyRoles(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userRoles, err := h.service.GetByUserID(userID.(uuid.UUID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRoles)
}

// @Summary Get user roles by role ID
// @Tags user-roles
// @Produce json
// @Param roleId path string true "Role ID"
// @Success 200 {array} entities.UserRole
// @Failure 400 {object} map[string]interface{}
// @Router /user-roles/role/{roleId} [get]
func (h *UserRoleHandler) GetByRoleID(c *gin.Context) {
	roleIDStr := c.Param("roleId")
	roleID, err := uuid.Parse(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	userRoles, err := h.service.GetByRoleID(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userRoles)
}
