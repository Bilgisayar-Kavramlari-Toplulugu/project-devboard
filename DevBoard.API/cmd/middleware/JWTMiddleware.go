package middleware

import (
	"net/http"
	"strings"

	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware - JWT token doğrulama middleware
func JWTMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization header'ını al
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header required",
			})
			c.Abort()
			return
		}

		// Bearer token formatını kontrol et
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authorization header format. Use: Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Token'ı doğrula
		claims, err := jwtService.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Context'e kullanıcı bilgilerini ekle
		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// OptionalJWTMiddleware - Opsiyonel JWT token doğrulama
// Token varsa doğrular ve context'e ekler, yoksa devam eder
func OptionalJWTMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.Next()
			return
		}

		tokenString := parts[1]
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
