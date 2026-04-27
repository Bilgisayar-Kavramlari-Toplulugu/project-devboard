package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupCountryRoutes(r *gin.RouterGroup, countryHandler *handler.CountryHandler, jwtService services.JWTService, config *config.Config) {
	countryAuthenticatedRoute := r.Group("/countries")
	countryAuthenticatedRoute.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))

	countryAuthenticatedRoute.GET("/", countryHandler.GetAll)
	countryAuthenticatedRoute.GET("/alphabetical", countryHandler.GetAllAlphabetical)
	countryAuthenticatedRoute.GET("/:id", countryHandler.GetById)

	countryAuthorizedRoute := countryAuthenticatedRoute.Group("/")
	countryAuthorizedRoute.Use(middleware.AuthorizationMiddleware("admin"))
	countryAuthorizedRoute.POST("/", countryHandler.Create)
	countryAuthorizedRoute.PUT("/:id", countryHandler.Update)
	countryAuthorizedRoute.DELETE("/:id", countryHandler.Delete)
}
