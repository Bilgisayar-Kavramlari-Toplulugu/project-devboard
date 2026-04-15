package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"project-devboard/internal/config"
	"project-devboard/internal/domain/entities"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Claims - JWT token claims
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

// TokenPair - Access ve Refresh token çifti
type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// JWTService - JWT işlemleri için interface
type JWTService interface {
	GenerateTokenPair(userID uuid.UUID, role string) (*TokenPair, error)
	ValidateAccessToken(tokenString string) (*Claims, error)
	ValidateRefreshToken(tokenString string) (*Claims, error)
	RefreshTokens(refreshToken string, userID uuid.UUID) (*TokenPair, error)
	CreateSession(userID uuid.UUID, refreshToken string) error
	RevokeSession(refreshTokenHash string) error
}

type jwtService struct {
	cfg *config.Config
	db  *gorm.DB
}

// NewJWTService - JWTService oluştur
func NewJWTService(cfg *config.Config, db *gorm.DB) JWTService {
	return &jwtService{
		cfg: cfg,
		db:  db,
	}
}

// GenerateTokenPair - Access ve Refresh token çifti oluştur
func (s *jwtService) GenerateTokenPair(userID uuid.UUID, role string) (*TokenPair, error) {
	// Access token oluştur
	accessTokenExpiry := time.Now().Add(time.Duration(s.cfg.AccessTokenExpireHours) * time.Hour)
	accessClaims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID.String(),
			Issuer:    "saivier-api",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	// Refresh token oluştur
	refreshTokenExpiry := time.Now().Add(time.Duration(s.cfg.RefreshTokenExpireHours) * time.Hour)
	refreshClaims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshTokenExpiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID.String(),
			Issuer:    "saivier-api-refresh",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    accessTokenExpiry,
	}, nil
}

// ValidateAccessToken - Access token'ı doğrula
func (s *jwtService) ValidateAccessToken(tokenString string) (*Claims, error) {
	return s.validateToken(tokenString, "saivier-api")
}

// ValidateRefreshToken - Refresh token'ı doğrula
func (s *jwtService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return s.validateToken(tokenString, "saivier-api-refresh")
}

// validateToken - Token'ı doğrula (internal)
func (s *jwtService) validateToken(tokenString, expectedIssuer string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Issuer kontrolü
	if claims.Issuer != expectedIssuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

// RefreshTokens - Refresh token ile yeni token pair oluştur
func (s *jwtService) RefreshTokens(refreshToken string, userID uuid.UUID) (*TokenPair, error) {
	// Refresh token'ı doğrula
	claims, err := s.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Session kontrolü ve kullanıcı bilgilerini al - refresh token hash'i ile ara
	hash := sha256.Sum256([]byte(refreshToken))
	refreshTokenHash := hex.EncodeToString(hash[:])

	var user entities.User
	err = s.db.Preload("UserRoles.Role").Where("refresh_token_hash = ? AND id = ?", refreshTokenHash, claims.UserID).First(&user).Error
	if err != nil {
		return nil, errors.New("session not found or already revoked")
	}

	// Session süresi kontrolü
	if user.RefreshTokenExp == nil || user.RefreshTokenExp.Before(time.Now()) {
		return nil, errors.New("session expired")
	}

	role := "Developer"
	if len(user.UserRoles) > 0 && user.UserRoles[0].Role.Name != "" {
		role = user.UserRoles[0].Role.Name
	}

	// Yeni token pair oluştur
	return s.GenerateTokenPair(claims.UserID, role)
}

// CreateSession - Yeni session oluştur
func (s *jwtService) CreateSession(userID uuid.UUID, refreshToken string) error {
	// Refresh token'ı hash'le
	hash := sha256.Sum256([]byte(refreshToken))
	refreshTokenHash := hex.EncodeToString(hash[:])

	expiresAt := time.Now().Add(time.Duration(s.cfg.RefreshTokenExpireHours) * time.Hour)

	return s.db.Model(&entities.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"refresh_token_hash": refreshTokenHash,
		"refresh_token_exp":  expiresAt,
	}).Error
}

// RevokeSession - Session'ı iptal et
func (s *jwtService) RevokeSession(refreshTokenHash string) error {
	return s.db.Model(&entities.User{}).
		Where("refresh_token_hash = ?", refreshTokenHash).
		Updates(map[string]interface{}{
			"refresh_token_hash": nil,
			"refresh_token_exp":  nil,
		}).Error
}
