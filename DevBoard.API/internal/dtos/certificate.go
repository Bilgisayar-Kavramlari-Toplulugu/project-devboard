package dtos

import "time"

type CertificateDTO struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	Degree     *string    `json:"degree,omitempty"`
	IssueDate  time.Time  `json:"issueDate"`
	ExpireDate *time.Time `json:"expireDate,omitempty"`
	Url        *string    `json:"url,omitempty"`
}
