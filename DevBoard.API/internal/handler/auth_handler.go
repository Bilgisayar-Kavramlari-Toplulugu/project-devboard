package handler

import (
	"net/http"

	"project-devboard/internal/config"
	"project-devboard/internal/dtos"
	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"
	"project-devboard/pkg/response"
	"project-devboard/pkg/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService services.AuthService
	validator   *validator.Validator
	cfg         *config.Config
}

func NewAuthHandler(authService services.AuthService, validator *validator.Validator, cfg *config.Config) *AuthHandler {
	return &AuthHandler{authService: authService, validator: validator, cfg: cfg}
}

// Login godoc
// @Summary      User login
// @Description  Login with email and password, sets session cookies, and returns a success message
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.LoginRequest  true  "Login Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Invalid credentials"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dtos.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}

	_, tokenPair, err := h.authService.Login(req)
	if err != nil {
		c.Error(err)
		return
	}

	setAuthCookies(c, h.cfg, tokenPair)
	response.Message(c, http.StatusOK, "Logged in successfully")
}

// Signup godoc
// @Summary      User signup
// @Description  Register a new user with email and password, sets session cookies, and returns a success message
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.SignupRequest  true  "Signup Request"
// @Success      201      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      409      {object}  response.Response  "Email already exists"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req dtos.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}

	userID := userIDFromContext(c, uuid.Nil)

	tokenPair, err := h.authService.Signup(req, userID)
	if err != nil {
		c.Error(err)
		return
	}

	setAuthCookies(c, h.cfg, tokenPair)
	response.Message(c, http.StatusCreated, "Signed up successfully")
}

// Refresh godoc
// @Summary      Refresh tokens
// @Description  Rotate session cookies using a valid refresh token and return a success message
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.RefreshTokenRequest  true  "Refresh Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Invalid refresh token"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken := readRefreshToken(c, h.cfg)
	if refreshToken == "" {
		var req dtos.RefreshTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
			return
		}
		if err := h.validator.Validate(req); err != nil {
			c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
			return
		}
		refreshToken = req.RefreshToken
	}

	tokenPair, err := h.authService.Refresh(refreshToken)
	if err != nil {
		clearAuthCookies(c, h.cfg)
		c.Error(err)
		return
	}

	setAuthCookies(c, h.cfg, tokenPair)
	response.Message(c, http.StatusOK, "Session refreshed successfully")
}

// Logout godoc
// @Summary      User logout
// @Description  Revoke the refresh token and end the session
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.LogoutRequest  true  "Logout Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken := readRefreshToken(c, h.cfg)
	if refreshToken == "" && requestHasBody(c) {
		var req dtos.LogoutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
			return
		}
		if err := h.validator.Validate(req); err != nil {
			c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
			return
		}
		refreshToken = req.RefreshToken
	}

	if refreshToken != "" {
		if err := h.authService.Logout(refreshToken); err != nil {
			c.Error(err)
			return
		}
	}

	clearAuthCookies(c, h.cfg)
	response.Message(c, http.StatusOK, "Logged out successfully")
}

// GetMe godoc
// @Summary      Get current user
// @Description  Get the currently authenticated user's information using the session cookie
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  UserEnvelope
// @Failure      401  {object}  response.Response  "Unauthorized"
// @Failure      404  {object}  response.Response  "User not found"
// @Failure      500  {object}  response.Response  "Internal Server Error"
// @Router       /auth/me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	userID := userIDFromContext(c, uuid.Nil)
	if userID == uuid.Nil {
		c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
		return
	}

	user, err := h.authService.GetMe(userID)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, http.StatusOK, dtos.NewUserResponse(user))
}

// ForgotPassword godoc
// @Summary      Request password reset
// @Description  Send a password reset email to the user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.ForgotPasswordRequest  true  "Forgot Password Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var req dtos.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}

	if err := h.authService.ForgotPassword(req.Email); err != nil {
		c.Error(err)
		return
	}

	response.Message(c, http.StatusOK, "If the email exists, a password reset link has been sent")
}

// ResetPassword godoc
// @Summary      Reset password
// @Description  Reset password using the token received via email
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      dtos.ResetPasswordRequest  true  "Reset Password Request"
// @Success      200      {object}  MessageEnvelope
// @Failure      400      {object}  response.Response  "Bad Request"
// @Failure      401      {object}  response.Response  "Invalid or expired token"
// @Failure      500      {object}  response.Response  "Internal Server Error"
// @Router       /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req dtos.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(apperrors.New(apperrors.InvalidRequest, apperrors.ErrInvalidRequest))
		return
	}
	if err := h.validator.Validate(req); err != nil {
		c.Error(apperrors.Validation(toAppFieldErrors(h.validator.FormatErrors(err))))
		return
	}

	if err := h.authService.ResetPassword(req.Token, req.NewPassword); err != nil {
		c.Error(err)
		return
	}

	response.Message(c, http.StatusOK, "Password reset successfully. Please login with your new password.")
}
