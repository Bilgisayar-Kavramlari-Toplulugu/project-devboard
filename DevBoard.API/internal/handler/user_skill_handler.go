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

type UserSkillHandler struct {
	service   services.UserSkillService
	validator *validator.Validator
}

func NewUserSkillHandler(service services.UserSkillService, validator *validator.Validator) *UserSkillHandler {
	return &UserSkillHandler{
		service:   service,
		validator: validator,
	}
}

func (h *UserSkillHandler) GetUserSkills(c *gin.Context) {
	userID := userIDFromContext(c, uuid.Nil)
	if userID == uuid.Nil {
		c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
		return
	}

	userSkills, err := h.service.GetUserSkills(userID)
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]dtos.UserSkillResponse, len(userSkills))
	for i, us := range userSkills {
		resp[i] = dtos.UserSkillResponse{
			SkillId:   us.Skill.Id,
			SkillName: us.Skill.Name,
		}
	}

	response.Success(c, http.StatusOK, resp)
}

func (h *UserSkillHandler) CreateUserSkills(c *gin.Context) {
	var req dtos.CreateUserSkillRequest
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
	if userID == uuid.Nil {
		c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
		return
	}

	userSkills, err := h.service.CreateUserSkills(userID, req.SkillIds)
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]dtos.UserSkillResponse, len(userSkills))
	for i, us := range userSkills {
		resp[i] = dtos.UserSkillResponse{
			SkillId:   us.Skill.Id,
			SkillName: us.Skill.Name,
		}
	}

	response.Success(c, http.StatusCreated, resp)
}

func (h *UserSkillHandler) UpdateUserSkills(c *gin.Context) {
	var req dtos.UpdateUserSkillRequest
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
	if userID == uuid.Nil {
		c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
		return
	}

	userSkills, err := h.service.UpdateUserSkills(userID, req.SkillIds)
	if err != nil {
		c.Error(err)
		return
	}

	resp := make([]dtos.UserSkillResponse, len(userSkills))
	for i, us := range userSkills {
		resp[i] = dtos.UserSkillResponse{
			SkillId:   us.Skill.Id,
			SkillName: us.Skill.Name,
		}
	}

	response.Success(c, http.StatusOK, resp)
}
