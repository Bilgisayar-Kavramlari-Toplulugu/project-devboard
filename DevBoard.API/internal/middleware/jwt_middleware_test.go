package middleware

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtServiceStub struct {
	validateAccessToken func(token string) (*services.Claims, error)
}

func (s jwtServiceStub) GenerateTokenPair(userID uuid.UUID, role string) (*services.TokenPair, error) {
	return nil, errors.New("not implemented")
}

func (s jwtServiceStub) ValidateAccessToken(tokenString string) (*services.Claims, error) {
	if s.validateAccessToken != nil {
		return s.validateAccessToken(tokenString)
	}
	return nil, errors.New("not implemented")
}

func (s jwtServiceStub) ValidateRefreshToken(tokenString string) (*services.Claims, error) {
	return nil, errors.New("not implemented")
}

func (s jwtServiceStub) RefreshTokens(refreshToken string, userID uuid.UUID) (*services.TokenPair, error) {
	return nil, errors.New("not implemented")
}

func (s jwtServiceStub) CreateSession(userID uuid.UUID, refreshToken string) error {
	return errors.New("not implemented")
}

func (s jwtServiceStub) RevokeSession(refreshTokenHash string) error {
	return errors.New("not implemented")
}

func TestJWTMiddlewareRejectsAuthorizationHeaderWithoutCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(GlobalErrorHandler())
	router.Use(JWTMiddleware(jwtServiceStub{}, "access_token"))
	router.GET("/protected", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer legacy-token")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestJWTMiddlewareAcceptsAccessTokenCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.Use(GlobalErrorHandler())
	router.Use(JWTMiddleware(jwtServiceStub{
		validateAccessToken: func(token string) (*services.Claims, error) {
			if token != "cookie-token" {
				return nil, errors.New("unexpected token")
			}

			return &services.Claims{
				UserID: uuid.New(),
				Role:   "Developer",
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
				},
			}, nil
		},
	}, "access_token"))
	router.GET("/protected", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "access_token", Value: "cookie-token"})

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}
