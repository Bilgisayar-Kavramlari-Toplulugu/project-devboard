package dtos

type SkillCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	SkillTypeId int    `json:"skill_type_id" validate:"required"`
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
