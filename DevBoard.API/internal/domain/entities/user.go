package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                 uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Username           string     `gorm:"type:varchar(500);uniqueIndex;not null"`
	Email              string     `gorm:"type:varchar(500);not null"`
	Password           string     `gorm:"type:varchar(500);not null" json:"-"`
	Firstname          string     `gorm:"type:varchar(500);not null"`
	Lastname           string     `gorm:"type:varchar(500);not null"`
	PhoneNumber        *string    `gorm:"type:varchar(500)"`
	CityId             *int       `gorm:"type:integer"`
	Birthdate          *time.Time `gorm:"type:timestamp"`
	Gender             *int       `gorm:"type:integer"`
	ProfilePicturePath *string    `gorm:"type:varchar(500)"`
	Title              *string    `gorm:"type:varchar(500)"`
	RefreshTokenHash   *string    `gorm:"type:varchar(500)" json:"-"`
	RefreshTokenExp    *time.Time `gorm:"type:timestamp" json:"-"`
	UserAgent          *string    `gorm:"type:varchar(500)" json:"-"`
	IsEmailValidated   bool       `gorm:"type:boolean;default:false"`
	BaseEntity

	City                    *City                  `gorm:"foreignKey:CityId"`
	UserRoles               []UserRole             `gorm:"foreignKey:UserId"`
	UserJobTypes            []UserJobType          `gorm:"foreignKey:UserId"`
	UserWorkLocationTypes   []UserWorkLocationType `gorm:"foreignKey:UserId"`
	UserSkills              []UserSkill            `gorm:"foreignKey:UserId"`
	Certificates            []Certificate          `gorm:"foreignKey:UserId"`
	Experiences             []Experience           `gorm:"foreignKey:UserId"`
	Educations              []Education            `gorm:"foreignKey:UserId"`
	ProfessionalProfiles    []ProfessionalProfile  `gorm:"foreignKey:UserId"`
	References              []Reference            `gorm:"foreignKey:UserId"`
	Projects                []Project              `gorm:"foreignKey:UserId"`
	SavedFilters            []SavedFilter          `gorm:"foreignKey:OwnerId"`
	SavedDevelopersByUser   []SavedDeveloper       `gorm:"foreignKey:UserId"`
	SavedDevelopersByDev    []SavedDeveloper       `gorm:"foreignKey:DeveloperId"`
	SavedProjects           []SavedProject         `gorm:"foreignKey:UserId"`
	SentMessages            []Message              `gorm:"foreignKey:SenderId"`
	ReceivedMessages        []Message              `gorm:"foreignKey:ReceiverId"`
	PublicEndorsementsSent  []PublicEndorsement    `gorm:"foreignKey:SenderUserId"`
	ProjectDevelopers       []ProjectDeveloper     `gorm:"foreignKey:DeveloperId"`
	ProjectEndorsementsSent []ProjectEndorsement   `gorm:"foreignKey:SenderId"`
}

func (User) TableName() string { return "Users" }
