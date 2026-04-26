package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupDeveloperDashboardRoutes(rg *gin.RouterGroup, h *handler.DeveloperDashboardHandler, jwtService services.JWTService, config *config.Config) {
	dashboard := rg.Group("/developer/dashboard")
	dashboard.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))
	dashboard.Use(middleware.AuthorizationMiddleware("Developer"))
	{
		dashboard.GET("/me", h.GetCurrentUserDashboardData)
		dashboard.GET("/:username", h.GetDashboardDataByUsername)
	}
}
