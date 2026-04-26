package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupCityRoutes(r *gin.RouterGroup, cityHandler *handler.CityHandler, jwtService services.JWTService, config *config.Config) {
	cityAuthenticatedRoute := r.Group("/cities")
	cityAuthenticatedRoute.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))

	cityAuthenticatedRoute.GET("/", cityHandler.GetAll)
	cityAuthenticatedRoute.GET("/:id", cityHandler.GetById)
	cityAuthenticatedRoute.GET("/country/:countryId", cityHandler.GetByCountryId)

	cityAuthorizedRoute := cityAuthenticatedRoute.Group("/")
	cityAuthorizedRoute.Use(middleware.AuthorizationMiddleware("admin"))
	cityAuthorizedRoute.POST("/", cityHandler.Create)
	cityAuthorizedRoute.PUT("/:id", cityHandler.Update)
	cityAuthorizedRoute.DELETE("/:id", cityHandler.Delete)
}