package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupUserSkillRoutes(r *gin.RouterGroup, handler *handler.UserSkillHandler, jwtService services.JWTService, cfg *config.Config) {
	userSkills := r.Group("/user-skills")
	userSkills.Use(middleware.JWTMiddleware(jwtService, cfg.AccessTokenCookieName))
	userSkills.GET("", handler.GetUserSkills)
	userSkills.POST("", handler.CreateUserSkills)
	userSkills.PUT("", handler.UpdateUserSkills)
}
