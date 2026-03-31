package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"time"

	"project-devboard/internal/config"
	domain "project-devboard/internal/domain/entities"
	"project-devboard/internal/services"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	jwtService services.JWTService
	cfg        *config.Config
	db         *gorm.DB
	validator  *validator.Validator
}

func NewAuthHandler(jwtService services.JWTService, cfg *config.Config, db *gorm.DB, validator *validator.Validator) *AuthHandler {
	return &AuthHandler{
		jwtService: jwtService,
		cfg:        cfg,
		db:         db,
		validator:  validator,
	}
}

// LoginRequest - Login isteği
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RefreshRequest - Token yenileme isteği
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// LogoutRequest - Logout isteği
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// ForgotPasswordRequest - Şifremi unuttum isteği
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest - Şifre sıfırlama isteği
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

// SignupRequest - Kullanıcı kaydı isteği
type SignupRequest struct {
	Email       string  `json:"email" validate:"required,email"`
	Password    string  `json:"password" validate:"required,min=8"`
	Firstname   string  `json:"firstname" validate:"required,min=2"`
	Lastname    string  `json:"lastname" validate:"required,min=2"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"` // E.164 format: +[country_code][number]
}

// Login godoc
// @Summary      User login
// @Description  Login with email and password, returns access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login Request"
// @Success      200  {object}  services.TokenPair "Token Pair"
// @Failure      400  {object}  response.Response "Bad Request"
// @Failure      401  {object}  response.Response "Invalid credentials"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Get user with roles and subscriptions
	var user domain.User
	if err := h.db.Preload("UserRoles.Role").Where("email = ?", req.Email).First(&user).Error; err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Get role (default: Basic)
	role := "Basic"
	if len(user.UserRoles) > 0 && user.UserRoles[0].Role.Name != "" {
		role = user.UserRoles[0].Role.Name
	}

	// Generate tokens
	tokenPair, err := h.jwtService.GenerateTokenPair(user.ID, role)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate tokens")
		return
	}

	// Device info oluştur
	deviceInfo := map[string]string{
		"ip":         c.ClientIP(),
		"user_agent": c.Request.UserAgent(),
	}
	deviceInfoBytes, _ := json.Marshal(deviceInfo)

	// Session oluştur
	err = h.jwtService.CreateSession(
		user.ID,
		tokenPair.RefreshToken,
		deviceInfoBytes,
		c.ClientIP(),
		c.Request.UserAgent(),
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create session")
		return
	}

	response.Success(c, http.StatusOK, tokenPair)
}

// Signup godoc
// @Summary      User signup
// @Description  Register a new user with email and password, returns access and refresh tokens
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body SignupRequest true "Signup Request"
// @Success      201  {object}  services.TokenPair "Token Pair"
// @Failure      400  {object}  response.Response "Bad Request"
// @Failure      409  {object}  response.Response "Email already exists"
// @Router       /auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Email uniqueness check
	var existingUser domain.User
	err := h.db.Where("email = ?", req.Email).First(&existingUser).Error
	if err == nil {
		// User already exists
		response.Error(c, http.StatusConflict, "Email already exists")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Create new user
	newUser := domain.User{
		ID:               uuid.New(),
		Email:            req.Email,
		Password:         string(hashedPassword),
		Firstname:        req.Firstname,
		Lastname:         req.Lastname,
		PhoneNumber:      req.PhoneNumber,
		IsEmailValidated: false,
	}

	if err := h.db.Create(&newUser).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Create UserRole with Basic role
	basicRoleID := uuid.MustParse("84aab341-99e8-48be-9789-917af5175ed3")
	userRole := domain.UserRole{
		ID:     uuid.New(),
		UserID: newUser.ID,
		RoleID: basicRoleID,
	}

	if err := h.db.Create(&userRole).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to assign role")
		return
	}

	// Generate tokens
	role := "Basic" // New users have Basic role
	tokenPair, err := h.jwtService.GenerateTokenPair(newUser.ID, role)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate tokens")
		return
	}

	// Create device info
	deviceInfo := map[string]string{
		"ip":         c.ClientIP(),
		"user_agent": c.Request.UserAgent(),
	}
	deviceInfoBytes, _ := json.Marshal(deviceInfo)

	// Create session
	err = h.jwtService.CreateSession(
		newUser.ID,
		tokenPair.RefreshToken,
		deviceInfoBytes,
		c.ClientIP(),
		c.Request.UserAgent(),
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create session")
		return
	}

	response.Success(c, http.StatusCreated, tokenPair)
}

// Refresh godoc
// @Summary      Refresh tokens
// @Description  Get new access and refresh tokens using a valid refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RefreshRequest true "Refresh Request"
// @Success      200  {object}  services.TokenPair "New Token Pair"
// @Failure      400  {object}  response.Response "Bad Request"
// @Failure      401  {object}  response.Response "Invalid refresh token"
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Refresh token'ı doğrula
	claims, err := h.jwtService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid or expired refresh token")
		return
	}

	// Yeni token pair oluştur
	tokenPair, err := h.jwtService.RefreshTokens(req.RefreshToken, claims.UserID)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Eski session'ı revoke et
	oldHash := sha256.Sum256([]byte(req.RefreshToken))
	oldRefreshTokenHash := hex.EncodeToString(oldHash[:])
	_ = h.jwtService.RevokeSession(oldRefreshTokenHash)

	// Yeni session oluştur
	deviceInfo := map[string]string{
		"ip":         c.ClientIP(),
		"user_agent": c.Request.UserAgent(),
	}
	deviceInfoBytes, _ := json.Marshal(deviceInfo)

	err = h.jwtService.CreateSession(
		claims.UserID,
		tokenPair.RefreshToken,
		deviceInfoBytes,
		c.ClientIP(),
		c.Request.UserAgent(),
	)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create session")
		return
	}

	response.Success(c, http.StatusOK, tokenPair)
}

// Logout godoc
// @Summary      User logout
// @Description  Revoke the refresh token and end the session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LogoutRequest true "Logout Request"
// @Success      200  {object}  map[string]string "Success message"
// @Failure      400  {object}  response.Response "Bad Request"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Refresh token'ı hash'le
	hash := sha256.Sum256([]byte(req.RefreshToken))
	refreshTokenHash := hex.EncodeToString(hash[:])

	// Session'ı revoke et
	err := h.jwtService.RevokeSession(refreshTokenHash)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to logout")
		return
	}

	response.Success(c, http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// GetMe godoc
// @Summary      Get current user
// @Description  Get the currently authenticated user's information
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  entities.User "Current User"
// @Failure      401  {object}  response.Response "Unauthorized"
// @Router       /auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var user domain.User
	err := h.db.Preload("UserRoles.Role").First(&user, "id = ?", userID.(uuid.UUID)).Error
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	// Şifreyi gizle
	user.Password = ""

	response.Success(c, http.StatusOK, user)
}

// ForgotPassword godoc
// @Summary      Request password reset
// @Description  Send a password reset email to the user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body ForgotPasswordRequest true "Forgot Password Request"
// @Success      200  {object}  map[string]string "Success message"
// @Failure      400  {object}  response.Response "Bad Request"
// @Router       /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Kullanıcıyı bul (email var mı kontrol et ama hata mesajı verme - güvenlik)
	var user domain.User
	err := h.db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		// Güvenlik: Email bulunamasa bile başarılı mesajı dön
		response.Success(c, http.StatusOK, gin.H{"message": "If the email exists, a password reset link has been sent"})
		return
	}

	// Eski token'ları invalidate et (aynı anda tek valid token kuralı)
	h.db.Model(&domain.PasswordResetToken{}).
		Where("user_id = ? AND used_at IS NULL", user.ID).
		Update("used_at", time.Now())

	// Yeni token oluştur (32 byte random = 64 hex karakter)
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate reset token")
		return
	}
	rawToken := hex.EncodeToString(tokenBytes)

	// Token'ı hash'le
	hash := sha256.Sum256([]byte(rawToken))
	tokenHash := hex.EncodeToString(hash[:])

	// Database'e kaydet
	resetToken := &domain.PasswordResetToken{
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(time.Duration(h.cfg.PasswordResetExpireMinutes) * time.Minute),
		CreatedAt: time.Now(),
	}

	if err := h.db.Create(resetToken).Error; err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create reset token")
		return
	}

	// Email gönder
	// if err := h.emailService.SendPasswordResetEmail(user.Email, rawToken); err != nil {
	// 	response.Error(c, http.StatusInternalServerError, "Failed to send reset email")
	// 	return
	// }

	response.Success(c, http.StatusOK, gin.H{"message": "If the email exists, a password reset link has been sent"})
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Reset password using the token received via email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body ResetPasswordRequest true "Reset Password Request"
// @Success      200  {object}  map[string]string "Success message"
// @Failure      400  {object}  response.Response "Bad Request"
// @Failure      401  {object}  response.Response "Invalid or expired token"
// @Router       /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(c, err)
		return
	}

	// Token'ı hash'le
	hash := sha256.Sum256([]byte(req.Token))
	tokenHash := hex.EncodeToString(hash[:])

	// Token'ı bul
	var resetToken domain.PasswordResetToken
	err := h.db.Where("token_hash = ?", tokenHash).First(&resetToken).Error
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	// Token geçerli mi kontrol et
	if !resetToken.IsValid() {
		response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	// Şifreyi hash'le
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Transaction başlat
	tx := h.db.Begin()

	// Şifreyi güncelle
	if err := tx.Model(&domain.User{}).Where("id = ?", resetToken.UserID).Update("password", string(hashedPassword)).Error; err != nil {
		tx.Rollback()
		response.Error(c, http.StatusInternalServerError, "Failed to update password")
		return
	}

	// Token'ı kullanıldı olarak işaretle
	now := time.Now()
	if err := tx.Model(&resetToken).Update("used_at", &now).Error; err != nil {
		tx.Rollback()
		response.Error(c, http.StatusInternalServerError, "Failed to mark token as used")
		return
	}

	tx.Commit()

	response.Success(c, http.StatusOK, gin.H{"message": "Password reset successfully. Please login with your new password."})
}
