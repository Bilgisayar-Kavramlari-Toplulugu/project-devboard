package main

import (
	"log"
	"strings"

	"project-devboard/internal/config"
	"project-devboard/internal/db_plugins"
	"project-devboard/internal/handler"
	"project-devboard/internal/repository"
	"project-devboard/internal/routes"
	"project-devboard/internal/services"
	"project-devboard/pkg/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "project-devboard/docs" // yourprojectname yerine kendi module adınızı yazın

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
	// 1. Config
	cfg := config.Load()

	// 2. Database - önce veritabanını oluştur
	createDatabaseIfNotExists(cfg.DatabaseURL)

	// Sonra bağlan
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed:", err)
	}
	// Auto Migration
	config.RunMigrations(db)

	//Plugins
	db.Use(&db_plugins.TimestampPlugin{})

	// 3. Validator
	v := validator.New()

	// 4. Dependency Injection

	// Level-0 Entity Repositories
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)

	// Level-0 Entity Services
	userService := services.NewUserService(userRepo)
	roleService := services.NewRoleService(roleRepo)
	userRoleService := services.NewUserRoleService(userRoleRepo)

	// Level-0 Entity Handlers
	userHandler := handler.NewUserHandler(userService, v)
	roleHandler := handler.NewRoleHandler(roleService, v)
	userRoleHandler := handler.NewUserRoleHandler(userRoleService, v)

	// JWT Service, Email Service ve Auth Handler
	jwtService := services.NewJWTService(cfg, db)
	authHandler := handler.NewAuthHandler(jwtService, cfg, db, v)

	// Device Token Service ve Handler (legacy - will be refactored)

	// Metadata Service

	// 5. Router
	r := gin.Default()

	// Redis'i başlat
	// redisClient := config.InitRedis()
	// defer config.CloseRedis() // Program kapanırken kapat

	// Route yapılandırması Redisli
	// routeConfigWithRedis := &routes.RouteConfig{
	// 	DB:          db,
	// 	RedisClient: redisClient,
	// 	UserHandler: userHandler,
	// }
	// routes.SetupRoutes(r, routeConfigWithRedis)

	// Route yapılandırması No Redis
	routeConfig := &routes.RouteConfig{
		DB:              db,
		UserHandler:     userHandler,
		RoleHandler:     roleHandler,
		UserRoleHandler: userRoleHandler,
		AuthHandler:     authHandler,
		JWTService:      jwtService,
	}

	// 5. CORS Middleware - Flutter web'den gelen isteklere izin ver
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:*", "http://127.0.0.1:*", "http://localhost", "http://127.0.0.1"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Development: Allow all localhost and 127.0.0.1 origins
			return strings.HasPrefix(origin, "http://localhost") ||
				strings.HasPrefix(origin, "http://127.0.0.1")
		},
	}))

	// 6. Setup Routes
	routes.SetupRoutes(r, routeConfig)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 6. Start Server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal("Server start failed:", err)
	}
}

func createDatabaseIfNotExists(databaseURL string) {
	// postgres veritabanına bağlan (her zaman vardır)
	adminDB := strings.Replace(databaseURL, "dbname=SaivierDB", "dbname=postgres", 1)

	db, err := gorm.Open(postgres.Open(adminDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Admin DB Connection Failed:", err)
	}

	// SaivierDB var mı kontrol et
	result := db.Exec("SELECT 1 FROM pg_database WHERE datname = 'SaivierDB'")
	if result.RowsAffected == 0 {
		// Yoksa oluştur
		db.Exec("CREATE DATABASE \"SaivierDB\"")
		log.Println("Database SaivierDB created successfully")
	}

	// Admin bağlantısını kapat
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

// ## Ne Zaman Neyi Kullanmalı?

// | İhtiyaç | Kullan | Neden |
// |---------|--------|-------|
// | Rate limiting sayacı | **Redis** | Hızlı, TTL ile otomatik temizlik |
// | API call history | **PostgreSQL** | Kalıcı kayıt, analiz için |
// | Geçici sayaçlar | **Redis** | Performans |
// | Kalıcı veri | **PostgreSQL** | Güvenilirlik |

// ## Özet Akış
// ```
// İstek geldi
//   ↓
// 1. API key doğrula (PostgreSQL)
//   ↓
// 2. Rate limit kontrolü (Redis) ← HIZLI
//   ↓
// 3. İstek işlenir
//   ↓
// 4. API call log (PostgreSQL) ← ASYNC, bloklama yok
