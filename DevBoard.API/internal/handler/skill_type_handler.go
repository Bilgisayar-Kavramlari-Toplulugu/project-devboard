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

type SkillTypeHandler struct {
	service   services.SkillTypeService
	validator *validator.Validator
}

func NewSkillTypeHandler(service services.SkillTypeService, validator *validator.Validator) *SkillTypeHandler {
	return &SkillTypeHandler{service: service, validator: validator}
}

func (h *SkillTypeHandler) Create(c *gin.Context) {
	var req dtos.SkillTypeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	userID := userIDFromContext(c, uuid.New())
	skillTypeId, err := h.service.CreateSkillType(req.Name, userID)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.SkillTypeResponse{Id: skillTypeId, Name: req.Name})
}

func (h *SkillTypeHandler) GetAll(c *gin.Context) {
	skillTypes, err := h.service.GetAllSkillTypes()
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, skillTypes)
}
func (h *SkillTypeHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	skillTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	skillType, err := h.service.GetSkillTypeById(skillTypeId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, skillType)
}

func (h *SkillTypeHandler) Update(c *gin.Context) {
	id := c.Param("id")

	skillTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	var req dtos.SkillTypeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}

	err = h.service.UpdateSkillType(skillTypeId, req.Name)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.SkillTypeResponse{Id: skillTypeId, Name: req.Name})
}
func (h *SkillTypeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	skillTypeId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	err = h.service.DeleteSkillType(skillTypeId)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusNoContent, nil)
}
