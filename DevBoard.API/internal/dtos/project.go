package dtos

import "time"

type ProjectDTO struct {
	Id               int64      `json:"id"`
	TypeId           int        `json:"typeId"`
	TypeName         string     `json:"typeName"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"shortDescription"`
	Description      string     `json:"description"`
	StartDate        time.Time  `json:"startDate"`
	EndDate          *time.Time `json:"endDate,omitempty"`
	Url              *string    `json:"url,omitempty"`
}
