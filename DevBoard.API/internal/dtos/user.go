package dtos

import (
	"time"

	"project-devboard/internal/domain/entities"

	"github.com/google/uuid"
)

type UserCreateRequest struct {
	Username    string  `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=6"`
	Firstname   string  `json:"firstname" validate:"required,min=2,max=100"`
	Lastname    string  `json:"lastname" validate:"required,min=2,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type UserUpdateRequest struct {
	Email       *string `json:"email" validate:"omitempty,email"`
	Password    *string `json:"password" validate:"omitempty,min=6"`
	Firstname   *string `json:"firstname" validate:"omitempty,min=2,max=100"`
	Lastname    *string `json:"lastname" validate:"omitempty,min=2,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type UserResponse struct {
	ID               uuid.UUID  `json:"id"`
	Username         string     `json:"username"`
	Email            string     `json:"email"`
	Firstname        string     `json:"firstname"`
	Lastname         string     `json:"lastname"`
	PhoneNumber      *string    `json:"phone_number,omitempty"`
	CityID           *int       `json:"city_id,omitempty"`
	Birthdate        *time.Time `json:"birthdate,omitempty"`
	Gender           *int       `json:"gender,omitempty"`
	ProfilePicture   *string    `json:"profile_picture,omitempty"`
	Title            *string    `json:"title,omitempty"`
	IsEmailValidated bool       `json:"is_email_validated"`
	IsActive         bool       `json:"is_active"`
	Roles            []string   `json:"roles,omitempty"`
	CreatedOn        time.Time  `json:"created_on"`
	LastModifiedOn   time.Time  `json:"last_modified_on"`
}

func NewUserResponse(user *entities.User) *UserResponse {
	if user == nil {
		return nil
	}

	roles := make([]string, 0, len(user.UserRoles))
	for _, userRole := range user.UserRoles {
		if userRole.Role.Name != "" {
			roles = append(roles, userRole.Role.Name)
		}
	}

	return &UserResponse{
		ID:               user.Id,
		Username:         user.Username,
		Email:            user.Email,
		Firstname:        user.Firstname,
		Lastname:         user.Lastname,
		PhoneNumber:      user.PhoneNumber,
		CityID:           user.CityId,
		Birthdate:        user.Birthdate,
		Gender:           user.Gender,
		ProfilePicture:   user.ProfilePicturePath,
		Title:            user.Title,
		IsEmailValidated: user.IsEmailValidated,
		IsActive:         user.BaseEntity.IsActive,
		Roles:            roles,
		CreatedOn:        user.BaseEntity.CreatedOn,
		LastModifiedOn:   user.BaseEntity.LastModifiedOn,
	}
}

func NewUserResponses(users []entities.User) []UserResponse {
	result := make([]UserResponse, 0, len(users))
	for i := range users {
		if user := NewUserResponse(&users[i]); user != nil {
			result = append(result, *user)
		}
	}

	return result
}
