package dtos

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
