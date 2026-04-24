package handler

import (
	"net/http"
	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkLocationTypeHandler struct {
	service   services.WorkLocationTypeService
	validator *validator.Validator
}

func NewWorkLocationTypeHandler(service services.WorkLocationTypeService, validator *validator.Validator) *WorkLocationTypeHandler {
	return &WorkLocationTypeHandler{service: service, validator: validator}
}

// @Summary Create a new work location type
// @Description Create a new work location type with the given name (admin only)
// @Tags WorkLocationTypes
// @Accept json
// @Produce json
// @Param request body dtos.WorkLocationTypeCreateRequest true "Work location type details"
// @Success 201 {object} response.Response{data=dtos.WorkLocationTypeResponse}
// @Failure 400 {object} response.Response
// @Failure 409 {object} response.Response
// @Security BearerAuth
// @Router /work-location-types [post]
func (h *WorkLocationTypeHandler) Create(c *gin.Context) {
	var req dtos.WorkLocationTypeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}
	userID := userIDFromContext(c, uuid.Nil)
	id, err := h.service.CreateWorkLocationType(req.Name, userID)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.WorkLocationTypeResponse{Id: id, Name: req.Name})
}

// @Summary Get all work location types
// @Description Get a paginated list of all work location types
// @Tags WorkLocationTypes
// @Produce json
// @Param limit query int false "Number of results per page (max 100)" default(20)
// @Param page query int false "Page number" default(1)
// @Success 200 {object} response.Response{data=[]dtos.WorkLocationTypeResponse}
// @Security BearerAuth
// @Router /work-location-types [get]
func (h *WorkLocationTypeHandler) GetAll(c *gin.Context) {
	limit := 20
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			if v > 100 {
				v = 100
			}
			limit = v
		}
	}
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 1 {
			offset = (v - 1) * limit
		}
	}
	list, err := h.service.GetAllWorkLocationTypes(limit, offset)
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.WorkLocationTypeResponse, len(list))
	for i, wlt := range list {
		result[i] = dtos.WorkLocationTypeResponse{Id: wlt.Id, Name: wlt.Name}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get a work location type by ID
// @Description Get details of a work location type by its ID
// @Tags WorkLocationTypes
// @Produce json
// @Param id path int true "Work Location Type ID"
// @Success 200 {object} response.Response{data=dtos.WorkLocationTypeResponse}
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /work-location-types/{id} [get]
func (h *WorkLocationTypeHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	wlt, err := h.service.GetWorkLocationTypeById(id)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.WorkLocationTypeResponse{Id: wlt.Id, Name: wlt.Name})
}

// @Summary Update a work location type
// @Description Update an existing work location type's name (admin only)
// @Tags WorkLocationTypes
// @Accept json
// @Produce json
// @Param id path int true "Work Location Type ID"
// @Param request body dtos.WorkLocationTypeUpdateRequest true "Updated details"
// @Success 200 {object} response.Response{data=dtos.WorkLocationTypeResponse}
// @Failure 404 {object} response.Response
// @Failure 409 {object} response.Response
// @Security BearerAuth
// @Router /work-location-types/{id} [put]
func (h *WorkLocationTypeHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	var req dtos.WorkLocationTypeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}
	if err := h.service.UpdateWorkLocationType(id, req.Name); err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.WorkLocationTypeResponse{Id: id, Name: req.Name})
}

// @Summary Delete a work location type
// @Description Delete a work location type by its ID (admin only)
// @Tags WorkLocationTypes
// @Param id path int true "Work Location Type ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /work-location-types/{id} [delete]
func (h *WorkLocationTypeHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.service.DeleteWorkLocationType(id); err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
