package dtos

import (
	"project-devboard/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type DeveloperDashboardResponse struct {
	Id                 uuid.UUID               `json:"id"`
	Email              string                  `json:"email"`
	Firstname          string                  `json:"firstname"`
	Lastname           string                  `json:"lastname"`
	PhoneNumber        *string                 `json:"phoneNumber"`
	Birthdate          *time.Time              `json:"birthdate"`
	Gender             *int                    `json:"gender"`
	ProfilePicturePath *string                 `json:"profilePicturePath"`
	Title              *string                 `json:"title"`
	City               *entities.City          `json:"city"`
	Skills             []entities.UserSkill    `json:"skills"`
	Certificates       []entities.Certificate  `json:"certificates"`
	Experiences        []entities.Experience   `json:"experiences"`
	Educations         []entities.Education    `json:"educations"`
	ProfessionalProfiles []entities.ProfessionalProfile `json:"professionalProfiles"`
	JobTypes           []entities.UserJobType  `json:"jobTypes"`
	WorkLocationTypes  []entities.UserWorkLocationType `json:"workLocationTypes"`
	Projects           []entities.Project      `json:"projects"`
	SentMessages       []entities.Message      `json:"sentMessages"`
	ReceivedMessages   []entities.Message      `json:"receivedMessages"`
	PublicEndorsementsSent []entities.PublicEndorsement `json:"publicEndorsementsSent"`
	ProjectEndorsementsSent []entities.ProjectEndorsement `json:"projectEndorsementsSent"`
	References         []entities.Reference    `json:"references"`
}

func NewDeveloperDashboardResponse(user *entities.User) DeveloperDashboardResponse {
	return DeveloperDashboardResponse{
		Id:                 user.Id,
		Email:              user.Email,
		Firstname:          user.Firstname,
		Lastname:           user.Lastname,
		PhoneNumber:        user.PhoneNumber,
		Birthdate:          user.Birthdate,
		Gender:             user.Gender,
		ProfilePicturePath: user.ProfilePicturePath,
		Title:              user.Title,
		City:               user.City,
		Skills:             user.UserSkills,
		Certificates:       user.Certificates,
		Experiences:        user.Experiences,
		Educations:         user.Educations,
		ProfessionalProfiles: user.ProfessionalProfiles,
		JobTypes:           user.UserJobTypes,
		WorkLocationTypes:  user.UserWorkLocationTypes,
		Projects:           user.Projects,
		SentMessages:       user.SentMessages,
		ReceivedMessages:   user.ReceivedMessages,
		PublicEndorsementsSent: user.PublicEndorsementsSent,
		ProjectEndorsementsSent: user.ProjectEndorsementsSent,
		References:         user.References,
	}
}
