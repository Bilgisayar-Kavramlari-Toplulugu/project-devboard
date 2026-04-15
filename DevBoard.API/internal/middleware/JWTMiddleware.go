package middleware

import (
	"strings"

	"project-devboard/internal/services"
	"project-devboard/pkg/apperrors"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware - JWT token doğrulama middleware
func JWTMiddleware(jwtService services.JWTService, accessTokenCookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractAccessToken(c, accessTokenCookieName)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		claims, err := jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			c.Error(apperrors.New(apperrors.InvalidToken, apperrors.ErrInvalidToken))
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// OptionalJWTMiddleware - Opsiyonel JWT token doğrulama
// Token varsa doğrular ve context'e ekler, yoksa devam eder
func OptionalJWTMiddleware(jwtService services.JWTService, accessTokenCookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractOptionalAccessToken(c, accessTokenCookieName)
		if err != nil || tokenString == "" {
			c.Next()
			return
		}

		claims, err := jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func extractAccessToken(c *gin.Context, accessTokenCookieName string) (string, error) {
	if token := tokenFromCookie(c, accessTokenCookieName); token != "" {
		return token, nil
	}

	return "", apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized)
}

func extractOptionalAccessToken(c *gin.Context, accessTokenCookieName string) (string, error) {
	if token := tokenFromCookie(c, accessTokenCookieName); token != "" {
		return token, nil
	}

	return "", nil
}

func tokenFromCookie(c *gin.Context, cookieName string) string {
	if cookieName == "" {
		return ""
	}

	token, err := c.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(token)
}
