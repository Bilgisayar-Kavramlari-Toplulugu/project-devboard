package routes

import (
	"project-devboard/internal/handler"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupUserRoleRoutes - UserRole endpoints
func SetupUserRoleRoutes(api *gin.RouterGroup, h *handler.UserRoleHandler, jwtService services.JWTService) {
	userroles := api.Group("/user-roles")
	{
		userroles.POST("", h.CreateUserRole)
		userroles.GET("", h.ListUserRoles)
		userroles.GET("/:id", h.GetUserRole)
		userroles.PUT("/:id", h.UpdateUserRole)
		userroles.DELETE("/:id", h.DeleteUserRole)
	}

	// Current user's roles
	api.GET("/user-roles/me", h.GetMyRoles)
	api.GET("/user-roles/role/:roleId", h.GetByRoleID)
}
