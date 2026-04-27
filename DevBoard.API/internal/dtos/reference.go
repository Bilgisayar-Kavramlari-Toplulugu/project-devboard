package dtos

type ReferenceDTO struct {
	Id           int64   `json:"id"`
	Firstname    string  `json:"firstname"`
	Lastname     string  `json:"lastname"`
	PhoneNumber  *string `json:"phoneNumber,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
}
