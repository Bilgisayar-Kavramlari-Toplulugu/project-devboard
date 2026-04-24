package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupSkillRoutes(api *gin.RouterGroup, skillHandler *handler.SkillHandler, jwtService services.JWTService, cfg *config.Config) {
	skillAuthenticatedRoute := api.Group("/skills")
	skillAuthenticatedRoute.Use(middleware.JWTMiddleware(jwtService, cfg.AccessTokenCookieName))

	skillAuthenticatedRoute.GET("/", skillHandler.GetAll)
	skillAuthenticatedRoute.GET("/:id", skillHandler.GetById)

	skillAuthorizedRoute := skillAuthenticatedRoute.Group("/")
	skillAuthorizedRoute.Use(middleware.AuthorizationMiddleware("admin"))

	skillAuthorizedRoute.POST("/", skillHandler.Create)
	skillAuthorizedRoute.DELETE("/:id", skillHandler.Delete)
	skillAuthorizedRoute.PUT("/:id", skillHandler.Update)
}
