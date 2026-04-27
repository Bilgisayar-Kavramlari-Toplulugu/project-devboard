package dtos

type JobTypeDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type JobTypeCreateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}

type JobTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JobTypeUpdateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}
