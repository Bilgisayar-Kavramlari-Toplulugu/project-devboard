package dtos

type CityCreateRequest struct {
	Name         string `json:"name" validate:"required,min=2,max=500"`
	Code         string `json:"code" validate:"max=500"`
	DisplayOrder *int   `json:"display_order"`
	CountryId    *int   `json:"country_id"`
}

type CityUpdateRequest struct {
	Name         string `json:"name" validate:"required,min=2,max=500"`
	Code         string `json:"code" validate:"max=500"`
	DisplayOrder *int   `json:"display_order"`
	CountryId    *int   `json:"country_id"`
}

type CityResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	DisplayOrder *int   `json:"display_order"`
	CountryId    *int   `json:"country_id"`
}