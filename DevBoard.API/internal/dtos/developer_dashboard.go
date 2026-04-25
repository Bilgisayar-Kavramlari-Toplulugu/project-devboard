package dtos

import (
	"project-devboard/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)

type DeveloperDashboardResponse struct {
	Id                      uuid.UUID                       `json:"id"`
	Email                   string                          `json:"email"`
	Firstname               string                          `json:"firstname"`
	Lastname                string                          `json:"lastname"`
	PhoneNumber             *string                         `json:"phoneNumber"`
	Birthdate               *time.Time                      `json:"birthdate"`
	Gender                  *int                            `json:"gender"`
	ProfilePicturePath      *string                         `json:"profilePicturePath"`
	Title                   *string                         `json:"title"`
	City                    *entities.City                  `json:"city"`
	Skills                  []entities.UserSkill            `json:"skills"`
	Certificates            []entities.Certificate          `json:"certificates"`
	Experiences             []entities.Experience           `json:"experiences"`
	Educations              []entities.Education            `json:"educations"`
	ProfessionalProfiles    []entities.ProfessionalProfile  `json:"professionalProfiles"`
	JobTypes                []entities.UserJobType          `json:"jobTypes"`
	WorkLocationTypes       []entities.UserWorkLocationType `json:"workLocationTypes"`
	Projects                []entities.Project              `json:"projects"`
	SentMessages            []entities.Message              `json:"sentMessages"`
	ReceivedMessages        []entities.Message              `json:"receivedMessages"`
	PublicEndorsementsSent  []entities.PublicEndorsement    `json:"publicEndorsementsSent"`
	ProjectEndorsementsSent []entities.ProjectEndorsement   `json:"projectEndorsementsSent"`
	References              []entities.Reference            `json:"references"`
}

func NewDeveloperDashboardResponse(user *entities.User) DeveloperDashboardResponse {
	return DeveloperDashboardResponse{
		Id:                      user.Id,
		Email:                   user.Email,
		Firstname:               user.Firstname,
		Lastname:                user.Lastname,
		PhoneNumber:             user.PhoneNumber,
		Birthdate:               user.Birthdate,
		Gender:                  user.Gender,
		ProfilePicturePath:      user.ProfilePicturePath,
		Title:                   user.Title,
		City:                    user.City,
		Skills:                  user.UserSkills,
		Certificates:            user.Certificates,
		Experiences:             user.Experiences,
		Educations:              user.Educations,
		ProfessionalProfiles:    user.ProfessionalProfiles,
		JobTypes:                user.UserJobTypes,
		WorkLocationTypes:       user.UserWorkLocationTypes,
		Projects:                user.Projects,
		SentMessages:            user.SentMessages,
		ReceivedMessages:        user.ReceivedMessages,
		PublicEndorsementsSent:  user.PublicEndorsementsSent,
		ProjectEndorsementsSent: user.ProjectEndorsementsSent,
		References:              user.References,
	}
}

type CurrentUserDashboardResponse struct {
	Id                      uuid.UUID                       `json:"id"`
	Email                   string                          `json:"email"`
	Firstname               string                          `json:"firstname"`
	Lastname                string                          `json:"lastname"`
	PhoneNumber             *string                         `json:"phoneNumber"`
	Birthdate               *time.Time                      `json:"birthdate"`
	Gender                  *int                            `json:"gender"`
	ProfilePicturePath      *string                         `json:"profilePicturePath"`
	Title                   *string                         `json:"title"`
	City                    *entities.City                  `json:"city"`
	Skills                  []entities.UserSkill            `json:"skills"`
	Certificates            []entities.Certificate          `json:"certificates"`
	Experiences             []entities.Experience           `json:"experiences"`
	Educations              []entities.Education            `json:"educations"`
	ProfessionalProfiles    []entities.ProfessionalProfile  `json:"professionalProfiles"`
	JobTypes                []entities.UserJobType          `json:"jobTypes"`
	WorkLocationTypes       []entities.UserWorkLocationType `json:"workLocationTypes"`
	Projects                []entities.Project              `json:"projects"`
	SentMessages            []entities.Message              `json:"sentMessages"`
	ReceivedMessages        []entities.Message              `json:"receivedMessages"`
	PublicEndorsementsSent  []entities.PublicEndorsement    `json:"publicEndorsementsSent"`
	ProjectEndorsementsSent []entities.ProjectEndorsement   `json:"projectEndorsementsSent"`
	References              []entities.Reference            `json:"references"`

	SkillsCount                  int     `json:"skillsCount"`
	CertificatesCount            int     `json:"certificatesCount"`
	ExperiencesCount             int     `json:"experiencesCount"`
	EducationsCount              int     `json:"educationsCount"`
	ProfessionalProfilesCount    int     `json:"professionalProfilesCount"`
	JobTypesCount                int     `json:"jobTypesCount"`
	WorkLocationTypesCount       int     `json:"workLocationTypesCount"`
	ProjectsCount                int     `json:"projectsCount"`
	SentMessagesCount            int     `json:"sentMessagesCount"`
	ReceivedMessagesCount        int     `json:"receivedMessagesCount"`
	PublicEndorsementsSentCount  int     `json:"publicEndorsementsSentCount"`
	ProjectEndorsementsSentCount int     `json:"projectEndorsementsSentCount"`
	ReferencesCount              int     `json:"referencesCount"`
	ProfileCompletePercentage    float64 `json:"profileCompletePercentage"`
}

func NewCurrentUserDashboardResponse(user *entities.User) CurrentUserDashboardResponse {
	return CurrentUserDashboardResponse{
		Id:                           user.Id,
		Email:                        user.Email,
		Firstname:                    user.Firstname,
		Lastname:                     user.Lastname,
		PhoneNumber:                  user.PhoneNumber,
		Birthdate:                    user.Birthdate,
		Gender:                       user.Gender,
		ProfilePicturePath:           user.ProfilePicturePath,
		Title:                        user.Title,
		City:                         user.City,
		Skills:                       user.UserSkills,
		Certificates:                 user.Certificates,
		Experiences:                  user.Experiences,
		Educations:                   user.Educations,
		ProfessionalProfiles:         user.ProfessionalProfiles,
		JobTypes:                     user.UserJobTypes,
		WorkLocationTypes:            user.UserWorkLocationTypes,
		Projects:                     user.Projects,
		SentMessages:                 user.SentMessages,
		ReceivedMessages:             user.ReceivedMessages,
		PublicEndorsementsSent:       user.PublicEndorsementsSent,
		ProjectEndorsementsSent:      user.ProjectEndorsementsSent,
		References:                   user.References,
		SkillsCount:                  len(user.UserSkills),
		CertificatesCount:            len(user.Certificates),
		ExperiencesCount:             len(user.Experiences),
		EducationsCount:              len(user.Educations),
		ProfessionalProfilesCount:    len(user.ProfessionalProfiles),
		JobTypesCount:                len(user.UserJobTypes),
		WorkLocationTypesCount:       len(user.UserWorkLocationTypes),
		ProjectsCount:                len(user.Projects),
		SentMessagesCount:            len(user.SentMessages),
		ReceivedMessagesCount:        len(user.ReceivedMessages),
		PublicEndorsementsSentCount:  len(user.PublicEndorsementsSent),
		ProjectEndorsementsSentCount: len(user.ProjectEndorsementsSent),
		ReferencesCount:              len(user.References),
		ProfileCompletePercentage:    calculateProfileCompletePercentage(user),
	}
}

func calculateProfileCompletePercentage(user *entities.User) float64 {
	percentage := 0.0

	// Base fields
	fields := map[string]bool{
		"firstname":         user.Firstname != "",
		"lastname":          user.Lastname != "",
		"city":              user.City != nil,
		"title":             user.Title != nil,
		"jobTypes":          len(user.UserJobTypes) > 0,
		"workLocationTypes": len(user.UserWorkLocationTypes) > 0,
	}

	// Profile picture
	if user.ProfilePicturePath != nil {
		fields["profilePicture"] = true
	}

	// Optional fields (counts)
	optionalFields := map[string]int{
		//"bio":      user.BioLength,
		"skills":       len(user.UserSkills),
		"experiences":  len(user.Experiences),
		"education":    len(user.Educations),
		"certificates": len(user.Certificates),
		"projects":     len(user.Projects),
		"references":   len(user.References),
		"messages":     len(user.SentMessages) + len(user.ReceivedMessages),
		"endorsements": len(user.PublicEndorsementsSent) + len(user.ProjectEndorsementsSent),
	}

	// Calculate percentage
	filledFields := 0
	for _, filled := range fields {
		if filled {
			filledFields++
		}
	}

	// Base field percentage (40% max)
	if len(fields) > 0 {
		percentage += (float64(filledFields) / float64(len(fields))) * 40.0
	}

	// Optional fields percentage (60% max)
	if len(optionalFields) > 0 {
		// Each optional field contributes equally
		// But we can give more weight to more important fields
		// For now, let's give each optional field 3-6% weight
		// Total optional weight = 60%
		// Let's spread it across the fields
		// Skills, experience, education, projects are most important

		weights := map[string]float64{
			"skills":       6.0,
			"experience":   6.0,
			"education":    6.0,
			"projects":     6.0,
			"references":   4.0,
			"messages":     4.0,
			"endorsements": 4.0,
			"bio":          4.0,
			"certificates": 4.0,
		}

		optionalPercentage := 0.0
		for field, weight := range weights {
			// Check if field has content
			var hasContent bool
			switch field {
			case "skills":
				hasContent = optionalFields["skills"] > 0
			case "experience":
				hasContent = optionalFields["experiences"] > 0
			case "education":
				hasContent = optionalFields["education"] > 0
			case "projects":
				hasContent = optionalFields["projects"] > 0
			case "references":
				hasContent = optionalFields["references"] > 0
			case "messages":
				hasContent = optionalFields["messages"] > 0
			case "endorsements":
				hasContent = optionalFields["endorsements"] > 0
			case "bio":
				hasContent = optionalFields["bio"] > 0
			case "certificates":
				hasContent = optionalFields["certificates"] > 0
			}

			if hasContent {
				optionalPercentage += weight
			}
		}

		// Cap at 60%
		if optionalPercentage > 60.0 {
			optionalPercentage = 60.0
		}

		percentage += optionalPercentage
	}

	// Round to nearest integer
	return float64(int(percentage + 0.5))
}
