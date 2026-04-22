package main

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"project-devboard/internal/config"
	"project-devboard/internal/db_plugins"
	"project-devboard/internal/handler"
	"project-devboard/internal/middleware"
	"project-devboard/internal/repository"
	"project-devboard/internal/routes"
	"project-devboard/internal/services"
	"project-devboard/pkg/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "project-devboard/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Dökümanı
// @version         1.0
// @description     Bu API'nin açıklaması
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Destek
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	cfg := config.Load()

	createDatabaseIfNotExists(cfg.DatabaseURL)

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}

	config.RunMigrations(db)
	db.Use(&db_plugins.TimestampPlugin{})

	v := validator.New()

	// Repositories
	userRepo := repository.NewUserRepository(db)
	skillTypeRepo := repository.NewSkillTypeRepository(db)
	jobTypeRepo := repository.NewJobTypeRepository(db)

	// Services
	jwtService := services.NewJWTService(cfg, db)
	authService := services.NewAuthService(db, userRepo, jwtService, cfg)
	userService := services.NewUserService(userRepo)
	skillTypeService := services.NewSkillTypeService(skillTypeRepo)
	jobTypeService := services.NewJobTypeService(jobTypeRepo)

	// Handlers
	authHandler := handler.NewAuthHandler(authService, v, cfg)
	userHandler := handler.NewUserHandler(userService, v)
	skillTypeHandler := handler.NewSkillTypeHandler(skillTypeService, v)
	jobTypeHandler := handler.NewJobTypeHandler(jobTypeService, v)

	// Router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(middleware.GlobalErrorHandler())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSAllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			for _, allowedOrigin := range cfg.CORSAllowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}

			return strings.HasPrefix(origin, "http://localhost") ||
				strings.HasPrefix(origin, "http://127.0.0.1")
		},
	}))

	routes.SetupRoutes(r, &routes.RouteConfig{
		DB:               db,
		UserHandler:      userHandler,
		AuthHandler:      authHandler,
		JWTService:       jwtService,
		SkillTypeHandler: skillTypeHandler,
		JobTypeHandler:   jobTypeHandler,
		Config:           cfg,
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal("Server start failed:", err)
	}
}

func createDatabaseIfNotExists(databaseURL string) {
	databaseName, err := extractDatabaseName(databaseURL)
	if err != nil {
		log.Fatal("Database name could not be determined from DATABASE_URL:", err)
	}

	adminDB := replaceDatabaseName(databaseURL, "postgres")

	db, err := gorm.Open(postgres.Open(adminDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Admin DB Connection Failed:", err)
	}

	var databaseCount int64
	if err := db.Raw("SELECT COUNT(1) FROM pg_database WHERE datname = ?", databaseName).Scan(&databaseCount).Error; err != nil {
		log.Fatal("Database existence check failed:", err)
	}

	if databaseCount == 0 {
		if err := db.Exec(`CREATE DATABASE "` + strings.ReplaceAll(databaseName, `"`, `""`) + `"`).Error; err != nil {
			log.Fatal("Database creation failed:", err)
		}
		log.Printf("Database %s created successfully\n", databaseName)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func extractDatabaseName(databaseURL string) (string, error) {
	if databaseURL == "" {
		return "", errors.New("empty DATABASE_URL")
	}

	if strings.HasPrefix(databaseURL, "postgres://") || strings.HasPrefix(databaseURL, "postgresql://") {
		parsed, err := url.Parse(databaseURL)
		if err != nil {
			return "", err
		}

		databaseName := strings.TrimPrefix(parsed.Path, "/")
		if databaseName == "" {
			return "", errors.New("missing database name in URL path")
		}

		return databaseName, nil
	}

	for _, part := range strings.Fields(databaseURL) {
		if strings.HasPrefix(part, "dbname=") {
			databaseName := strings.Trim(strings.TrimPrefix(part, "dbname="), `"'`)
			if databaseName == "" {
				return "", errors.New("empty dbname value")
			}

			return databaseName, nil
		}
	}

	return "", errors.New("missing dbname in DATABASE_URL")
}

func replaceDatabaseName(databaseURL, databaseName string) string {
	if strings.HasPrefix(databaseURL, "postgres://") || strings.HasPrefix(databaseURL, "postgresql://") {
		parsed, err := url.Parse(databaseURL)
		if err != nil {
			return databaseURL
		}

		parsed.Path = "/" + databaseName
		return parsed.String()
	}

	parts := strings.Fields(databaseURL)
	for i, part := range parts {
		if strings.HasPrefix(part, "dbname=") {
			parts[i] = "dbname=" + databaseName
			return strings.Join(parts, " ")
		}
	}

	return databaseURL
}
