package dtos

type UserSkillResponse struct {
	SkillId   int    `json:"skill_id"`
	SkillName string `json:"skill_name"`
}

type UserSkillDTO struct {
	Id        int64  `json:"id"`
	SkillId   int    `json:"skillId"`
	SkillName string `json:"skillName"`
}

type UserSkillQuery struct {
	Id int `uri:"id" binding:"required"`
}

type CreateUserSkillRequest struct {
	SkillIds []int `json:"skillIds" validate:"required,min=1,dive,required"`
}

type UpdateUserSkillRequest struct {
	SkillIds []int `json:"skillIds" validate:"required,min=1,dive,required"`
}

type DeleteUserSkillRequest struct {
	SkillIds []int `json:"skillIds" validate:"required,min=1,dive,required"`
}
