package routes

import (
	"project-devboard/internal/config"
	"project-devboard/internal/handler"
	"project-devboard/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouteConfig - Tüm route'lar için gerekli bağımlılıklar NoRedis
type RouteConfig struct {
	DB               *gorm.DB
	UserHandler      *handler.UserHandler
	AuthHandler      *handler.AuthHandler
	SkillTypeHandler *handler.SkillTypeHandler
	JWTService       services.JWTService
	Config           *config.Config
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
	SetupAuthRoutes(api, cfg.AuthHandler, cfg.JWTService, cfg.Config)

	SetupUserRoutes(api, cfg.UserHandler, cfg.JWTService, cfg.Config)

	SetupSkillTypeRoutes(api, cfg.SkillTypeHandler, cfg.JWTService, cfg.Config)
}
