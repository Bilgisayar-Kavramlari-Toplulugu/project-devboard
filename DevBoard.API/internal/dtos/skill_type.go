package dtos

import "project-devboard/internal/domain/entities"

type SkillTypeCreateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}

type SkillTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SkillTypeUpdateQuery struct {
	Id int `uri:"id" binding:"required"`
}

type SkillTypeUpdateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}

func NewSkillTypeResponse(skillType *entities.SkillType) *SkillTypeResponse {
	if skillType == nil {
		return nil
	}
	return &SkillTypeResponse{
		Id:   skillType.Id,
		Name: skillType.Name,
	}
}

func NewSkillTypeResponses(skillTypes []entities.SkillType) []SkillTypeResponse {
	result := make([]SkillTypeResponse, 0, len(skillTypes))
	for i := range skillTypes {
		if skillType := NewSkillTypeResponse(&skillTypes[i]); skillType != nil {
			result = append(result, *skillType)
		}
	}
	return result
}
