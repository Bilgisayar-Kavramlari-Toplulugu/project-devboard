package dtos

type EducationDTO struct {
	Id       int64        `json:"id"`
	Name     string       `json:"name"`
	Location *LocationDTO `json:"location"`
	GDPR     *float64     `json:"gdpr,omitempty"`
}
