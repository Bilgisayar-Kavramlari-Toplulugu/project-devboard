package dtos

type CountryCreateRequest struct {
	Name        string `json:"name" validate:"required,min=2,max=500"`
	FlagCode    string `json:"flag_code" validate:"required,max=500"`
	ShortCode   string `json:"short_code" validate:"required,max=500"`
	PhonePrefix string `json:"phone_prefix" validate:"required,max=500"`
}
type CountryCreateResponse struct {
	CountryDto
}
type CountryUpdateRequest struct {
	Name        string `json:"name" validate:"omitempty,min=2,max=500"`
	FlagCode    string `json:"flag_code" validate:"omitempty,max=500"`
	ShortCode   string `json:"short_code" validate:"omitempty,max=500"`
	PhonePrefix string `json:"phone_prefix" validate:"omitempty,max=500"`
}
type CountryResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	FlagCode    string `json:"flag_code"`
	ShortCode   string `json:"short_code"`
	PhonePrefix string `json:"phone_prefix"`
}
type CountryDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	FlagCode    string `json:"flag_code"`
	ShortCode   string `json:"short_code"`
	PhonePrefix string `json:"phone_prefix"`
}

