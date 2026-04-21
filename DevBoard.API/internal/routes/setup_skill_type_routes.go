package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupSkillTypeRoutes(r *gin.RouterGroup, skillTypeHandler *handler.SkillTypeHandler, jwtService services.JWTService, config *config.Config) {
	skillTypeAuthenticatedRoute := r.Group("/skill-types")
	skillTypeAuthenticatedRoute.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))

	skillTypeAuthenticatedRoute.GET("/", skillTypeHandler.GetAll)
	skillTypeAuthenticatedRoute.GET("/:id", skillTypeHandler.GetById)

	skillTypeAuthorizedRoute := skillTypeAuthenticatedRoute.Group("/")
	skillTypeAuthorizedRoute.Use(middleware.AuthorizationMiddleware("admin"))
	skillTypeAuthorizedRoute.POST("/", skillTypeHandler.Create)
	skillTypeAuthorizedRoute.PUT("/:id", skillTypeHandler.Update)
	skillTypeAuthorizedRoute.DELETE("/:id", skillTypeHandler.Delete)
}
