package dtos

import (
	"project-devboard/internal/domain/entities"
	"time"

	"github.com/google/uuid"
)



type DeveloperDashboardResponse struct {
	Id                      uuid.UUID                `json:"id"`
	Email                   string                   `json:"email"`
	Firstname               string                   `json:"firstname"`
	Lastname                string                   `json:"lastname"`
	PhoneNumber             *string                  `json:"phoneNumber"`
	Birthdate               *time.Time               `json:"birthdate"`
	Gender                  *int                     `json:"gender"`
	ProfilePicturePath      *string                  `json:"profilePicturePath"`
	Title                   *string                  `json:"title"`
	Location                *LocationDTO             `json:"location"`
	Skills                  []UserSkillDTO           `json:"skills"`
	Certificates            []CertificateDTO         `json:"certificates"`
	Experiences             []ExperienceDTO          `json:"experiences"`
	Educations              []EducationDTO           `json:"educations"`
	ProfessionalProfiles    []ProfessionalProfileDTO `json:"professionalProfiles"`
	JobTypes                []JobTypeDTO             `json:"jobTypes"`
	WorkLocationTypes       []WorkLocationTypeDTO    `json:"workLocationTypes"`
	Projects                []ProjectDTO             `json:"projects"`
	SentMessages            []MessageDTO             `json:"sentMessages"`
	ReceivedMessages        []MessageDTO             `json:"receivedMessages"`
	PublicEndorsementsSent  []PublicEndorsementDTO   `json:"publicEndorsementsSent"`
	ProjectEndorsementsSent []ProjectEndorsementDTO  `json:"projectEndorsementsSent"`
	References              []ReferenceDTO           `json:"references"`
}

type CurrentUserDashboardResponse struct {
	DeveloperDashboardResponse
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

// Map functions
func mapLocation(city *entities.City) *LocationDTO {
	if city == nil {
		return nil
	}
	countryName := ""
	if city.Country != nil {
		countryName = city.Country.Name
	}
	locName := city.Name
	if countryName != "" {
		if locName != "" {
			locName = locName + ", " + countryName
		} else {
			locName = countryName
		}
	}
	return &LocationDTO{
		CityId:      city.Id,
		CityName:    city.Name,
		CountryName: countryName,
		Name:        locName,
	}
}

func NewDeveloperDashboardResponse(user *entities.User) DeveloperDashboardResponse {
	// Map Skills
	skills := make([]UserSkillDTO, len(user.UserSkills))
	for i, s := range user.UserSkills {
		skills[i] = UserSkillDTO{Id: s.Id, SkillId: s.SkillId, SkillName: s.Skill.Name}
	}
	// Map Certificates
	certs := make([]CertificateDTO, len(user.Certificates))
	for i, c := range user.Certificates {
		certs[i] = CertificateDTO{Id: c.Id, Name: c.Name, Degree: c.Degree, IssueDate: c.IssueDate, ExpireDate: c.ExpireDate, Url: c.Url}
	}
	// Map Experiences
	exps := make([]ExperienceDTO, len(user.Experiences))
	for i, e := range user.Experiences {
		exps[i] = ExperienceDTO{Id: e.Id, Name: e.Name, Location: mapLocation(&e.City), Startdate: e.Startdate, Enddate: e.Enddate, Position: e.Position, Information: e.Information}
	}
	// Map Educations
	edus := make([]EducationDTO, len(user.Educations))
	for i, e := range user.Educations {
		edus[i] = EducationDTO{Id: e.Id, Name: e.Name, Location: mapLocation(&e.City), GDPR: e.GDPR}
	}
	// Map Profiles
	profs := make([]ProfessionalProfileDTO, len(user.ProfessionalProfiles))
	for i, p := range user.ProfessionalProfiles {
		pName := ""
		if p.Platform.Name != nil {
			pName = *p.Platform.Name
		}
		profs[i] = ProfessionalProfileDTO{Id: p.Id, PlatformId: p.PlatformId, PlatformName: pName, Url: p.Url}
	}
	// Map Job Types
	jobTypes := make([]JobTypeDTO, len(user.UserJobTypes))
	for i, jt := range user.UserJobTypes {
		jobTypes[i] = JobTypeDTO{Id: jt.JobTypeId, Name: jt.JobType.Name}
	}
	// Map Work Location Types
	workLocs := make([]WorkLocationTypeDTO, len(user.UserWorkLocationTypes))
	for i, wl := range user.UserWorkLocationTypes {
		wlName := ""
		if wl.WorkLocationType.Name != nil {
			wlName = *wl.WorkLocationType.Name
		}
		workLocs[i] = WorkLocationTypeDTO{Id: wl.WorkLocationTypeId, Name: wlName}
	}
	// Map Projects
	projects := make([]ProjectDTO, len(user.Projects))
	for i, p := range user.Projects {
		projects[i] = ProjectDTO{Id: p.Id, TypeId: p.TypeId, TypeName: p.Type.Name, Name: p.Name, ShortDescription: p.ShortDescription, Description: p.Description, StartDate: p.StartDate, EndDate: p.EndDate, Url: p.Url}
	}
	// Map Messages
	sentMsgs := make([]MessageDTO, len(user.SentMessages))
	for i, m := range user.SentMessages {
		sentMsgs[i] = MessageDTO{
			Id:                   m.Id,
			SenderId:             m.SenderId,
			ReceiverId:           m.ReceiverId,
			SenderEmailAddress:   m.Sender.Email,
			ReceiverEmailAddress: m.Receiver.Email,
			Subject:              m.Subject,
			Body:                 m.Body,
			IsHtml:               m.IsHtml,
			SentDate:             m.CreatedOn,
		}
	}
	receivedMsgs := make([]MessageDTO, len(user.ReceivedMessages))
	for i, m := range user.ReceivedMessages {
		receivedMsgs[i] = MessageDTO{
			Id:                   m.Id,
			SenderId:             m.SenderId,
			ReceiverId:           m.ReceiverId,
			SenderEmailAddress:   m.Sender.Email,
			ReceiverEmailAddress: m.Receiver.Email,
			Subject:              m.Subject,
			Body:                 m.Body,
			IsHtml:               m.IsHtml,
			SentDate:             m.CreatedOn,
		}
	}
	// Map Endorsements
	pubEnd := make([]PublicEndorsementDTO, len(user.PublicEndorsementsSent))
	for i, e := range user.PublicEndorsementsSent {
		pubEnd[i] = PublicEndorsementDTO{
			Id:                  e.Id,
			UserSkillId:         e.UserSkillId,
			SenderUserId:        e.SenderUserId,
			SenderUserFirstName: e.SenderUser.Firstname,
			SenderUserLastName:  e.SenderUser.Lastname,
			SenderEmailAddress:  e.SenderUser.Email,
			CreatedDate:         e.CreatedOn,
		}
	}
	projEnd := make([]ProjectEndorsementDTO, len(user.ProjectEndorsementsSent))
	for i, e := range user.ProjectEndorsementsSent {
		endorsementableName := ""
		if e.Endorsementable.Project != nil {
			endorsementableName = e.Endorsementable.Project.Name
		}
		projEnd[i] = ProjectEndorsementDTO{
			Id:                  e.Id,
			EndorsementableId:   e.EndorsementableId,
			EndorsementableName: endorsementableName,
			SenderId:            e.SenderId,
		}
	}
	// Map References
	refs := make([]ReferenceDTO, len(user.References))
	for i, r := range user.References {
		refs[i] = ReferenceDTO{Id: r.Id, Firstname: r.Firstname, Lastname: r.Lastname, PhoneNumber: r.PhoneNumber, EmailAddress: r.EmailAddress}
	}

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
		Location:                mapLocation(user.City),
		Skills:                  skills,
		Certificates:            certs,
		Experiences:             exps,
		Educations:              edus,
		ProfessionalProfiles:    profs,
		JobTypes:                jobTypes,
		WorkLocationTypes:       workLocs,
		Projects:                projects,
		SentMessages:            sentMsgs,
		ReceivedMessages:        receivedMsgs,
		PublicEndorsementsSent:  pubEnd,
		ProjectEndorsementsSent: projEnd,
		References:              refs,
	}
}

func NewCurrentUserDashboardResponse(user *entities.User) CurrentUserDashboardResponse {
	return CurrentUserDashboardResponse{
		DeveloperDashboardResponse:   NewDeveloperDashboardResponse(user),
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
		"skills":       len(user.UserSkills),
		"experiences":  len(user.Experiences),
		"education":    len(user.Educations),
		"certificates": len(user.Certificates),
		"projects":     len(user.Projects),
		"references":   len(user.References),
		"messages":     len(user.SentMessages) + len(user.ReceivedMessages),
		"endorsements": len(user.PublicEndorsementsSent) + len(user.ProjectEndorsementsSent),
	}

	filledFields := 0
	for _, filled := range fields {
		if filled {
			filledFields++
		}
	}

	if len(fields) > 0 {
		percentage += (float64(filledFields) / float64(len(fields))) * 40.0
	}

	if len(optionalFields) > 0 {
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

		if optionalPercentage > 60.0 {
			optionalPercentage = 60.0
		}
		percentage += optionalPercentage
	}

	return float64(int(percentage + 0.5))
}
