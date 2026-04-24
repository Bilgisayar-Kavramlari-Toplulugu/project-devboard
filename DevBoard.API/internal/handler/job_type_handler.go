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

type JobTypeHandler struct {
	service   services.JobTypeService
	validator *validator.Validator
}

func NewJobTypeHandler(service services.JobTypeService, validator *validator.Validator) *JobTypeHandler {
	return &JobTypeHandler{service: service, validator: validator}
}

// @Summary Create a new job type
// @Description Create a new job type with the given name
// @Tags JobTypes
// @Accept json
// @Produce json
// @Param request body dtos.JobTypeCreateRequest true "Job type details"
// @Success 201 {object} response.Response{data=dtos.JobTypeResponse}
// @Failure 400 {object} response.Response
// @Security BearerAuth
// @Router /job-types [post]
func (h *JobTypeHandler) Create(c *gin.Context) {
	var req dtos.JobTypeCreateRequest
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
	jobTypeId, err := h.service.CreateJobType(req.Name, userID)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.JobTypeResponse{Id: jobTypeId, Name: req.Name})
}

// @Summary Get all job types
// @Description Get a list of all job types
// @Tags JobTypes
// @Produce json
// @Success 200 {object} response.Response{data=[]dtos.JobTypeResponse}
// @Router /job-types [get]
func (h *JobTypeHandler) GetAll(c *gin.Context) {
	limit := 20
	offset := 0
	if l := c.Query("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil && v > 0 {
			limit = v
		}
	}
	if p := c.Query("page"); p != "" {
		if v, err := strconv.Atoi(p); err == nil && v > 1 {
			offset = (v - 1) * limit
		}
	}
	jobTypes, err := h.service.GetAllJobTypes(limit, offset)
	if err != nil {
		c.Error(err)
		return
	}
	result := make([]dtos.JobTypeResponse, len(jobTypes))
	for i, jt := range jobTypes {
		result[i] = dtos.JobTypeResponse{Id: jt.Id, Name: jt.Name}
	}
	response.Success(c, http.StatusOK, result)
}

// @Summary Get a job type by ID
// @Description Get details of a job type by its ID
// @Tags JobTypes
// @Produce json
// @Param id path int true "Job Type ID"
// @Success 200 {object} response.Response{data=dtos.JobTypeResponse}
// @Failure 404 {object} response.Response
// @Router /job-types/{id} [get]
func (h *JobTypeHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	jobTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	jobType, err := h.service.GetJobTypeById(jobTypeId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.JobTypeResponse{Id: jobType.Id, Name: jobType.Name})
}

// @Summary Update a job type
// @Description Update an existing job type's name
// @Tags JobTypes
// @Accept json
// @Produce json
// @Param id path int true "Job Type ID"
// @Param request body dtos.JobTypeUpdateRequest true "Updated details"
// @Success 200 {object} response.Response{data=dtos.JobTypeResponse}
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /job-types/{id} [put]
func (h *JobTypeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	jobTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	var req dtos.JobTypeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}

	err = h.service.UpdateJobType(jobTypeId, req.Name)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.JobTypeResponse{Id: jobTypeId, Name: req.Name})
}

// @Summary Delete a job type
// @Description Delete a job type by its ID
// @Tags JobTypes
// @Param id path int true "Job Type ID"
// @Success 204 "No Content"
// @Failure 404 {object} response.Response
// @Security BearerAuth
// @Router /job-types/{id} [delete]
func (h *JobTypeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	jobTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	err = h.service.DeleteJobType(jobTypeId)
	if err != nil {
		c.Error(err)
		return
	}
	c.Status(http.StatusNoContent)
}
