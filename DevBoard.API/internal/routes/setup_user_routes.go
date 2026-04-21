package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler, jwtService services.JWTService, config *config.Config) {
	users := rg.Group("/users")
	users.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))
	{
		users.GET("", h.List)
		users.GET("/:id", h.GetByID)
		users.POST("", h.Create)
		users.PUT("/:id", h.Update)
		users.DELETE("/:id", h.Delete)
	}
}
