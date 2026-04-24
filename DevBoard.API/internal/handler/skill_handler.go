package handler

import (
	"net/http"
	"project-devboard/internal/domain/entities"
	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SkillHandler struct {
	Service   services.SkillService
	validator *validator.Validator
}

func NewSkillHandler(service services.SkillService, validator *validator.Validator) *SkillHandler {
	return &SkillHandler{Service: service, validator: validator}
}

// Create godoc
// @Summary      Create a new skill
// @Description  Create a new skill with the input payload
// @Tags         skills
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.SkillCreateRequest  true  "Create Skill Request"
// @Success      201      {object}  SkillEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Unauthorized"
// @Failure      409      {object}  response.Response  "Skill already exists"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /skills [post]
func (h *SkillHandler) Create(c *gin.Context) {
	var req dtos.SkillCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	userId := userIDFromContext(c, uuid.New())
	skillId, err := h.Service.CreateSkill(services.CreateSkillInput{
		Name:        req.Name,
		SkillTypeId: req.SkillTypeId,
		UserId:      userId,
	})
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusCreated, dtos.NewSkillResponse(&entities.Skill{Id: skillId, Name: req.Name, SkillTypeId: req.SkillTypeId}))
}

// GetById godoc
// @Summary      Get a skill by ID
// @Description  Get details of a specific skill by ID
// @Tags         skills
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Skill ID"
// @Success      200  {object}  SkillEnvelope
// @Failure      400  {object}  response.Response  "Invalid ID"
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      404  {object}  response.Response  "Skill not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /skills/{id} [get]
func (h *SkillHandler) GetById(c *gin.Context) {
	var query dtos.SkillIdQuery
	if err := c.ShouldBindUri(&query); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(query); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	skill, err := h.Service.GetSkillById(query.Id)
	if err != nil {
		c.Error(err)
		return
	}
	response.Success(c, http.StatusOK, dtos.NewSkillResponse(skill))
}

// GetAll godoc
// @Summary      Get all skills
// @Description  Get a list of all skills
// @Tags         skills
// @Accept       json
// @Produce      json
// @Success      200  {object}  SkillListEnvelope
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /skills [get]
func (h *SkillHandler) GetAll(c *gin.Context) {
	skills, err := h.Service.GetAllSkills()
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusOK, dtos.NewSkillResponses(skills))
}

// Update godoc
// @Summary      Update a skill
// @Description  Update a skill's details by ID
// @Tags         skills
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.SkillUpdateRequest  true  "Update Skill Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Unauthorized"
// @Failure      404      {object}  response.Response  "Skill not found"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /skills/{id} [put]
func (h *SkillHandler) Update(c *gin.Context) {
	var req dtos.SkillUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	err := h.Service.UpdateSkill(services.UpdateSkillInput{
		Id:          req.Id,
		Name:        req.Name,
		SkillTypeId: req.SkillTypeId,
	})
	if err != nil {
		c.Error(err)
		return
	}
	response.Message(c, http.StatusOK, "Skill updated successfully")
}

// Delete godoc
// @Summary      Delete a skill
// @Description  Delete a skill by ID
// @Tags         skills
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Skill ID"
// @Success      200  {object}  MessageEnvelope
// @Failure      400  {object}  response.Response  "Invalid ID"
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      404  {object}  response.Response  "Skill not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /skills/{id} [delete]
func (h *SkillHandler) Delete(c *gin.Context) {
	var query dtos.SkillIdQuery
	if err := c.ShouldBindUri(&query); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(query); err != nil {
		fieldErrors := h.validator.FormatErrors(err)
		c.Error(apperrors.Validation(toAppFieldErrors(fieldErrors)))
		return
	}
	err := h.Service.DeleteSkill(query.Id)
	if err != nil {
		c.Error(err)
		return
	}
	response.Message(c, http.StatusOK, "Skill deleted successfully")
}
