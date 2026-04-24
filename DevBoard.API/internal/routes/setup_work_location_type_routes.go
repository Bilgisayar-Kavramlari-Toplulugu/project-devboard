package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupWorkLocationTypeRoutes(r *gin.RouterGroup, h *handler.WorkLocationTypeHandler, jwtService services.JWTService, config *config.Config) {
	authenticated := r.Group("/work-location-types")
	authenticated.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))

	authenticated.GET("/", h.GetAll)
	authenticated.GET("/:id", h.GetById)

	adminOnly := authenticated.Group("/")
	adminOnly.Use(middleware.AuthorizationMiddleware("admin"))
	adminOnly.POST("/", h.Create)
	adminOnly.PUT("/:id", h.Update)
	adminOnly.DELETE("/:id", h.Delete)
}
