package handler

import (
	"net/http"
	"strconv"

	domain "project-devboard/internal/domain/entities"
	"project-devboard/internal/services"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleHandler struct {
	service   services.RoleService
	validator *validator.Validator
}

func NewRoleHandler(service services.RoleService, v *validator.Validator) *RoleHandler {
	return &RoleHandler{
		service:   service,
		validator: v,
	}
}

// CreateRoleRequest - Role oluşturma request
type CreateRoleRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// UpdateRoleRequest - Role güncelleme request
type UpdateRoleRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

// @Summary Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param role body CreateRoleRequest true "Role data"
// @Success 201 {object} entities.Role
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /roles [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := &domain.Role{
		Name: req.Name,
	}

	if err := h.service.CreateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// @Summary Get role by ID
// @Tags roles
// @Produce json
// @Param id path string true "Role ID"
// @Success 200 {object} entities.Role
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /roles/{id} [get]
func (h *RoleHandler) GetRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	role, err := h.service.GetRole(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Summary List all roles
// @Tags roles
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} entities.Role
// @Router /roles [get]
func (h *RoleHandler) ListRoles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	roles, err := h.service.ListRoles(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// @Summary Update role
// @Tags roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Param role body UpdateRoleRequest true "Role data"
// @Success 200 {object} entities.Role
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /roles/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := &domain.Role{
		ID:   id,
		Name: req.Name,
	}

	if err := h.service.UpdateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// @Summary Delete role
// @Tags roles
// @Param id path string true "Role ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /roles/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	if err := h.service.DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
