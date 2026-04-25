package dtos

type ProfessionalProfileDTO struct {
	Id           int    `json:"id"`
	PlatformId   int    `json:"platformId"`
	PlatformName string `json:"platformName"`
	Url          string `json:"url"`
}
