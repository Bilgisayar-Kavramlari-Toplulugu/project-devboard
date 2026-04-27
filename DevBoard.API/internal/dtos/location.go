package dtos

type LocationDTO struct {
	CityId      int    `json:"cityId"`
	CityName    string `json:"cityName"`
	CountryName string `json:"countryName"`
	Name        string `json:"name"` // Format: CityName, CountryName
}
