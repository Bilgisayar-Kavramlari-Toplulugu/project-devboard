package dtos

type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required,min=6"`
}

type SignupRequest struct {
	Username    string  `json:"username" validate:"required,min=3,max=50,alphanum"`
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=8"`
	Firstname   string  `json:"firstname" validate:"required,min=2,max=100"`
	Lastname    string  `json:"lastname" validate:"required,min=2,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
