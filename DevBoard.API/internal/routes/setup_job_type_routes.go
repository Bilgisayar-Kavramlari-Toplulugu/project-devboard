package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupJobTypeRoutes(r *gin.RouterGroup, jobTypeHandler *handler.JobTypeHandler, jwtService services.JWTService, config *config.Config) {
	jobTypeAuthenticatedRoute := r.Group("/job-types")
	jobTypeAuthenticatedRoute.Use(middleware.JWTMiddleware(jwtService, config.AccessTokenCookieName))

	jobTypeAuthenticatedRoute.GET("/", jobTypeHandler.GetAll)
	jobTypeAuthenticatedRoute.GET("/:id", jobTypeHandler.GetById)

	jobTypeAuthorizedRoute := jobTypeAuthenticatedRoute.Group("/")
	jobTypeAuthorizedRoute.Use(middleware.AuthorizationMiddleware("admin"))
	jobTypeAuthorizedRoute.POST("/", jobTypeHandler.Create)
	jobTypeAuthorizedRoute.PUT("/:id", jobTypeHandler.Update)
	jobTypeAuthorizedRoute.DELETE("/:id", jobTypeHandler.Delete)
}
