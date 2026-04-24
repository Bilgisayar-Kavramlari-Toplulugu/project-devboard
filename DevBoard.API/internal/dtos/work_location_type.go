package dtos

type WorkLocationTypeCreateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}

type WorkLocationTypeUpdateRequest struct {
	Name string `json:"name" validate:"required,min=2,max=500"`
}

type WorkLocationTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
