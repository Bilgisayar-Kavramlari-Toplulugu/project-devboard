package dtos

import "time"

type ExperienceDTO struct {
	Id          int64        `json:"id"`
	Name        string       `json:"name"`
	Location    *LocationDTO `json:"location"`
	Startdate   time.Time    `json:"startDate"`
	Enddate     *time.Time   `json:"endDate,omitempty"`
	Position    *string      `json:"position,omitempty"`
	Information *int64       `json:"information,omitempty"`
}
