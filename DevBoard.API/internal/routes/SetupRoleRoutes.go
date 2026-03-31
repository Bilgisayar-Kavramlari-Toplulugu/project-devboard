package routes

import (
	"project-devboard/internal/handler"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoleRoutes - Role endpoints
func SetupRoleRoutes(api *gin.RouterGroup, h *handler.RoleHandler, jwtService services.JWTService) {
	roles := api.Group("/roles")
	{
		roles.POST("", h.CreateRole)
		roles.GET("", h.ListRoles)
		roles.GET("/:id", h.GetRole)
		roles.PUT("/:id", h.UpdateRole)
		roles.DELETE("/:id", h.DeleteRole)
	}
}
