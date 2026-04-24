package dtos

import "project-devboard/internal/domain/entities"

type SkillCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	SkillTypeId int    `json:"skill_type_id" validate:"required"`
}

type SkillIdQuery struct {
	Id int `uri:"id" validate:"required"`
}

type SkillUpdateRequest struct {
	Id          int    `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	SkillTypeId int    `json:"skill_type_id" validate:"required"`
}

type SkillResponse struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	SkillTypeId int               `json:"skill_type_id"`
	SkillType   SkillTypeResponse `json:"skill_type"`
}

func NewSkillResponse(skill *entities.Skill) *SkillResponse {
	if skill == nil {
		return nil
	}

	return &SkillResponse{
		Id:          skill.Id,
		Name:        skill.Name,
		SkillTypeId: skill.SkillTypeId,
		SkillType: SkillTypeResponse{
			Id:   skill.SkillType.Id,
			Name: skill.SkillType.Name,
		},
	}
}

func NewSkillResponses(skills []entities.Skill) []SkillResponse {
	result := make([]SkillResponse, 0, len(skills))
	for i := range skills {
		if skill := NewSkillResponse(&skills[i]); skill != nil {
			result = append(result, *skill)
		}
	}

	return result
}
