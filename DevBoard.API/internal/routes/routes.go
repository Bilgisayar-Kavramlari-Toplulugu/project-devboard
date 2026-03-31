package routes

import (
	"project-devboard/internal/handler"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouteConfig - Tüm route'lar için gerekli bağımlılıklar NoRedis
type RouteConfig struct {
	DB              *gorm.DB
	UserHandler     *handler.UserHandler
	RoleHandler     *handler.RoleHandler
	AuthHandler     *handler.AuthHandler
	UserRoleHandler *handler.UserRoleHandler
	JWTService      services.JWTService
}

// SetupRoutes - Tüm API route'larını yapılandırır
func SetupRoutes(r *gin.Engine, cfg *RouteConfig) {
	// API v1 grubu
	api := r.Group("/api/v1")

	// Middleware'leri uygula
	api.Use(
	//middleware.APIKeyRequiredMiddleware(cfg.DB),
	//middleware.RateLimitMiddleware(cfg.DB),
	//middleware.APICallLogger(cfg.DB),
	)

	// Auth route'larını yapılandır (public + protected)
	SetupAuthRoutes(api, cfg.AuthHandler, cfg.JWTService)

	SetupUserRoutes(api, cfg.UserHandler, cfg.JWTService)
	SetupRoleRoutes(api, cfg.RoleHandler, cfg.JWTService)
	SetupUserRoleRoutes(api, cfg.UserRoleHandler, cfg.JWTService)
}
