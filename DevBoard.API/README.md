# DevBoard API

The RESTful API backend for DevBoard, a professional network built for engineers to showcase skills, explore others' work, and connect with developers worldwide.

## Overview

DevBoard API is a Go-based service that provides authentication, user management, skills management, and professional profile management endpoints. It's designed with security, scalability, and developer experience in mind.

## Technologies

- **Language**: Go 1.25
- **Database**: PostgreSQL (configured via `internal/config`)
- **Authentication**: JWT tokens (see `internal/services/jwt_service.go`)
- **HTTP Framework**: Standard Go `net/http` with custom middleware
- **API Documentation**: Swagger/OpenAPI (in `docs/`)

## Project Structure

```
DevBoard.API/
├── cmd/api/              # Application entrypoint
├── internal/
│   ├── config/          # Configuration & migrations
│   ├── domain/          # Domain entities
│   ├── dtos/            # Data transfer objects
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # HTTP middleware (auth, errors, etc.)
│   ├── repository/      # Data access layer
│   ├── services/        # Business logic
│   └── db_plugins/      # Database plugins
├── pkg/
│   ├── apperrors/       # Error types & codes
│   ├── common/          # Common utilities
│   ├── response/        # Response formatting
│   └── validator/       # Input validation
└── docs/                # Swagger documentation
```

## Development

### Prerequisites

- Go 1.25+
- PostgreSQL 12+
- Docker & Docker Compose (optional)

### Local Setup

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Configure environment:

   The `.env` file is encrypted as `.env.age` using [age](https://github.com/FiloSottile/age). To decrypt it, your SSH public key must be registered as a recipient in `encrypt-decrypt-env-file.sh`.

   Install `age` if you don't have it:
   ```bash
   # macOS
   brew install age

   # Windows
   winget install FiloSottile.age
   ```

   Then decrypt from the repo root (where `.env.age` lives):
   ```bash
   # Run from DevBoard.API/ directory
   ../encrypt-decrypt-env-file.sh decrypt
   ```

   This will generate `.env` using your `~/.ssh/id_ed25519` (or `~/.ssh/id_rsa`) key.

3. Run migrations:
   ```bash
   go run ./cmd/api
   # Migrations run automatically on startup
   ```

4. Start the server:
   ```bash
   go run ./cmd/api
   ```

The API will be available at `http://localhost:8080`

## Docker Setup

### Build Locally

```bash
docker build -t devboard-api:0.0.1 .
```

### Run Locally

```bash
docker run --rm --name api -p 8080:8080 devboard-api:0.0.1
```

API will be available at: `http://localhost:8080`

### Run with Custom Port

```bash
docker run --rm --name api -e PORT=3000 -p 3000:3000 devboard-api:0.0.1
```

## Security Features

The Docker image includes:
- **Multi-stage build**: Compiles Go binary in builder stage, final image contains only the compiled binary
- **Non-root user**: App runs as `appuser` (UID 1001), not root
- **Minimal base image**: Alpine 3.19 reduces attack surface
- **Health checks**: Container health monitoring enabled
- **Static binary**: `CGO_ENABLED=0` for portability and security
- **.dockerignore**: Excludes unnecessary files (git, tests, docs) from build context

## API Documentation

Swagger documentation is available at:
```
GET /swagger/index.html
```

Generate updated docs:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g cmd/api/main.go
```

## Available Endpoints

### Authentication
- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user

### Users
- `GET /users/:id` - Get user profile
- `PUT /users/:id` - Update user profile

### Skills
- `GET /skills` - List available skills
- `POST /skills` - Create skill type

### Professional Profile
- `GET /profile/:id` - Get professional profile
- `PUT /profile/:id` - Update professional profile

See `routes/` for complete routing setup.

## Error Handling

The API uses standardized error responses with error codes defined in `pkg/apperrors/`. All errors include:
- HTTP status code
- Error code (application-specific)
- Human-readable message
- Request ID for debugging

## Rate Limiting

Rate limiting middleware is enabled by default. Configure in `internal/middleware/rate_limit_middleware.go`.

## Database Migrations

Migrations are automatically applied on startup. Custom migrations can be added in `internal/config/migrations.go`.

## Testing

```bash
go test ./...
```

## Contributing

1. Create a feature branch
2. Make your changes
3. Ensure tests pass
4. Submit a pull request

## License

See [LICENSE](../LICENSE)
