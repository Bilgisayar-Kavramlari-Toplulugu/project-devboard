package routes

import (
	"project-devboard/cmd/middleware"
	"project-devboard/internal/handler"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupAuthRoutes - Auth route'larını yapılandır
func SetupAuthRoutes(r *gin.RouterGroup, authHandler *handler.AuthHandler, jwtService services.JWTService) {
	// Public auth endpoints (no authentication required)
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.Signup)
		auth.POST("/refresh", authHandler.Refresh)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)
	}

	// Protected auth endpoints (authentication required)
	authProtected := r.Group("/auth")
	authProtected.Use(middleware.JWTMiddleware(jwtService))
	{
		authProtected.GET("/me", authHandler.GetMe)
	}
}
