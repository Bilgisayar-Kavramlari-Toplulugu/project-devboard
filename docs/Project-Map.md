# Project Map

Bu doküman, repo içindeki ana dosya ve klasörlerin ne işe yaradığını hızlıca görmek için hazırlanmıştır.

## Repo Kökü

- `README.md`
  Repo için kısa giriş noktası. Zamanla kurulum ve linklerin burada toplanması gerekir.
- `LICENSE`
  Lisans dosyası.
- `.github/`
  GitHub tarafındaki issue template ve benzeri repo otomasyonları.
- `docs/`
  Repo seviyesindeki insan odaklı dokümanlar.
- `docs/API-Response-Contract.md`
  API'nin güncel response envelope ve global error handling sözleşmesini açıklar.
- `docs/Auth-Session-Contract.md`
  Cookie-only auth, CSRF ve CORS davranışını açıklar.
- `DevBoard.API/`
  Çalışan Go backend uygulaması.

## DevBoard.API Kökü

- `go.mod`, `go.sum`
  Go modül bağımlılıkları.
- `Makefile`
  Uygulamayı çalıştırma, build alma ve swagger üretme komutları.
- `.env.example`
  Local environment değişkenlerinin örnek hali.
- `docs/`
  Swagger tarafından üretilen API doküman çıktıları.

## Uygulama Başlangıcı

- `DevBoard.API/cmd/api/main.go`
  Uygulamanın giriş noktası.
  Config yüklenir, veritabanı bağlantısı açılır, migration çalışır, dependency wiring yapılır ve Gin router ayağa kalkar.

## Konfigürasyon ve DB

- `DevBoard.API/internal/config/config.go`
  Environment değişkenlerini okuyup uygulama config nesnesine çevirir.
- `DevBoard.API/internal/config/migrations.go`
  GORM `AutoMigrate` çağrılarını merkezi olarak yönetir.
  Hangi tabloların ayağa kalkacağı burada belirlenir.
- `DevBoard.API/internal/db_plugins/timestamp_plugin.go`
  GORM tarafındaki timestamp davranışlarını uygulayan plugin.

## Domain Modeli

- `DevBoard.API/internal/domain/entities/`
  Veritabanı modelleri burada tutulur.
  Genel kural: her dosya bir entity veya ilişki tablosunu temsil eder.

Aktif user/auth akışında öne çıkanlar:

- `DevBoard.API/internal/domain/entities/user.go`
  Kullanıcı modeli.
- `DevBoard.API/internal/domain/entities/role.go`
  Rol modeli.
- `DevBoard.API/internal/domain/entities/user_role.go`
  Kullanıcı-rol ilişki tablosu.
- `DevBoard.API/internal/domain/entities/passwordresettoken.go`
  Şifre sıfırlama token kaydı.
- `DevBoard.API/internal/domain/entities/base_entity.go`
  Ortak audit alanları.

Diğer entity dosyaları:

- `certificate.go`, `education.go`, `experience.go`, `project.go`, `skill.go` ve benzerleri
  Gelecekteki domain kapsamı için hazırlanmış modellerdir.
  Route/service katmanı bunların hepsini henüz aktif kullanmıyor.

## DTO Katmanı

- `DevBoard.API/internal/dtos/auth.go`
  Auth endpoint request payload’ları.
- `DevBoard.API/internal/dtos/user.go`
  User endpoint request payload’ları ve API’ye dönen `UserResponse` şekli.

## HTTP Katmanı

- `DevBoard.API/internal/routes/routes.go`
  Tüm route kurulumunun merkezi giriş noktası.
- `DevBoard.API/internal/routes/SetupAuthRoutes.go`
  `/api/v1/auth/*` endpoint’lerini bağlar.
- `DevBoard.API/internal/routes/SetupUserRoutes.go`
  `/api/v1/users/*` endpoint’lerini bağlar.

- `DevBoard.API/internal/handler/auth_handler.go`
  Auth isteklerini alır, validation yapar, service katmanını çağırır ve response döner.
- `DevBoard.API/internal/handler/auth_cookies.go`
  Auth cookie set/clear ve CSRF cookie üretim yardımcılarını içerir.
- `DevBoard.API/internal/handler/user_handler.go`
  User CRUD isteklerini işler.
- `DevBoard.API/internal/handler/common.go`
  Handler’lar için ortak yardımcılar.
  Pagination okuma ve context’ten actor id çekme burada.

## Business Logic

- `DevBoard.API/internal/services/auth_service.go`
  Login, signup, token refresh, logout, `me`, forgot/reset password gibi auth kuralları.
- `DevBoard.API/internal/services/jwt_service.go`
  Access/refresh token üretme, doğrulama ve session yönetimi.
- `DevBoard.API/internal/services/user_service.go`
  User CRUD akışı ve iş kuralları.

## Veri Erişimi

- `DevBoard.API/internal/repository/base_repository.go`
  Generic CRUD altyapısı.
- `DevBoard.API/internal/repository/user_repository.go`
  User modeline özel sorgular.
  Özellikle email ile kullanıcı bulma ve role preload burada.

## Middleware

- `DevBoard.API/internal/middleware/GlobalErrorHandler.go`
  Uygulamadaki `AppError` nesnelerini tek bir API error formatına çevirir.
- `DevBoard.API/internal/middleware/JWTMiddleware.go`
  Access token cookie'sini doğrular ve kullanıcı bilgisini request context’ine yazar.
- `DevBoard.API/internal/middleware/CSRFMiddleware.go`
  Session cookie taşıyan unsafe isteklerde CSRF cookie/header eşleşmesini zorunlu kılar.
- `DevBoard.API/internal/middleware/RateLimitMiddleware.go`
  Şu an aktif kullanılmayan, rate limit için ayrılmış middleware taslağı.

## Paylaşılan Paketler

- `DevBoard.API/pkg/apperrors/`
  Uygulamanın ortak hata kodları, error definition’ları ve `AppError` tipi.
- `DevBoard.API/pkg/response/response.go`
  API’nin standart success/error response helper’ları.
- `DevBoard.API/pkg/validator/validator.go`
  `go-playground/validator` etrafındaki küçük wrapper.

## Response ve Error Handling

- Başarılı response'lar `response.Success(...)` ve `response.Message(...)` ile döner.
- Hata durumunda handler ve auth middleware doğrudan JSON yazmaz; `c.Error(...)` ile hatayı pipeline'a bırakır.
- `DevBoard.API/internal/middleware/GlobalErrorHandler.go`
  `AppError` nesnelerini tek response formatına çevirir ve panic recovery işini de üstlenir.
- `DevBoard.API/pkg/apperrors/app_error.go`
  HTTP status kategorisi ile client'a dönen business error code ayrımını taşır.
- Ayrıntılı sözleşme için `docs/API-Response-Contract.md` belgesine bakılmalıdır.
- Auth/session ayrıntıları için `docs/Auth-Session-Contract.md` belgesine bakılmalıdır.

## Swagger Çıktısı

- `DevBoard.API/docs/docs.go`
- `DevBoard.API/docs/swagger.json`
- `DevBoard.API/docs/swagger.yaml`

Bu dosyalar generated çıktıdır.
Elle düzenlemek yerine `Makefile` içindeki swagger komutu ile yeniden üretmek gerekir.

## Kısa Akış Özeti

1. `main.go` config ve DB bağlantısını açar.
2. `routes/` HTTP endpoint’lerini tanımlar.
3. `handler/` request validation ve response işini yapar.
4. `services/` business logic taşır.
5. `repository/` veritabanına erişir.
6. `domain/entities/` GORM modellerini temsil eder.
7. `pkg/response` ve `pkg/apperrors` API sözleşmesini standartlaştırır.

Dosya hiyerarşisi

├── DevBoard.API
│   ├── Makefile
│   ├── bin
│   │   └── api
│   ├── cmd
│   │   └── api
│   │       └── main.go
│   ├── docs
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── config
│   │   │   ├── config.go
│   │   │   └── migrations.go
│   │   ├── db_plugins
│   │   │   └── timestamp_plugin.go
│   │   ├── domain
│   │   │   └── entities
│   │   │       ├── base_entity.go
│   │   │       ├── certificate.go
│   │   │       ├── city.go
│   │   │       ├── country.go
│   │   │       ├── education.go
│   │   │       ├── experience.go
│   │   │       ├── job_type.go
│   │   │       ├── message.go
│   │   │       ├── message_template.go
│   │   │       ├── message_template_type.go
│   │   │       ├── passwordresettoken.go
│   │   │       ├── professional_platform.go
│   │   │       ├── professional_profile.go
│   │   │       ├── project.go
│   │   │       ├── project_developer.go
│   │   │       ├── project_endorsement.go
│   │   │       ├── project_endorsementable.go
│   │   │       ├── project_role.go
│   │   │       ├── project_type.go
│   │   │       ├── public_endorsement.go
│   │   │       ├── reference.go
│   │   │       ├── role.go
│   │   │       ├── saved_developer.go
│   │   │       ├── saved_filter.go
│   │   │       ├── saved_project.go
│   │   │       ├── skill.go
│   │   │       ├── skill_type.go
│   │   │       ├── user.go
│   │   │       ├── user_job_type.go
│   │   │       ├── user_project_skill.go
│   │   │       ├── user_role.go
│   │   │       ├── user_skill.go
│   │   │       ├── user_work_location_type.go
│   │   │       └── work_location_type.go
│   │   ├── dtos
│   │   │   ├── auth-xml.go
│   │   │   ├── auth.go
│   │   │   ├── skill.go
│   │   │   ├── skill_type.go
│   │   │   └── user.go
│   │   ├── handler
│   │   │   ├── auth_cookies.go
│   │   │   ├── auth_handler.go
│   │   │   ├── common.go
│   │   │   ├── swagger_types.go
│   │   │   └── user_handler.go
│   │   ├── middleware
│   │   │   ├── GlobalErrorHandler.go
│   │   │   ├── JWTMiddleware.go
│   │   │   ├── RateLimitMiddleware.go
│   │   │   └── jwt_middleware_test.go
│   │   ├── repository
│   │   │   ├── base_repository.go
│   │   │   └── user_repository.go
│   │   ├── routes
│   │   │   ├── SetupAuthRoutes.go
│   │   │   ├── SetupUserRoutes.go
│   │   │   └── routes.go
│   │   └── services
│   │       ├── auth_service.go
│   │       ├── jwt_service.go
│   │       └── user_service.go
│   └── pkg
│       ├── apperrors
│       │   ├── app_error.go
│       │   ├── code.go
│       │   └── def.go
│       ├── response
│       │   └── response.go
│       └── validator
│           └── validator.go
├── LICENSE
├── README.md
└── docs
    ├── Development-Workflow.md
    ├── Project-Map.md
    └── Verified-Commits-Guide.md