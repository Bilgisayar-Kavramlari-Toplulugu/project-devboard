# Local Docker Compose

Runs the full DevBoard stack (PostgreSQL, API, Frontend) locally using Docker Compose.

## Prerequisites

- Docker Desktop
- `age` encryption tool (for decrypting `.env`)

## 1. Decrypt the .env file

The `.env` file is encrypted as `.env.age` using [age](https://github.com/FiloSottile/age) and lives at the repo root. 
Your SSH public key must be added to the https://github.com/settings/keys
Your SSH public key must be registered as a recipient in `encrypt-decrypt-env-file.sh`.

Install `age` if you don't have it:

```bash
# macOS
brew install age

# Windows
winget install FiloSottile.age
```

Decrypt from the repo root:

```bash
# Run from the repo root (where encrypt-decrypt-env-file.sh lives)
./encrypt-decrypt-env-file.sh decrypt
```

This generates `.env` at the repo root using your `~/.ssh/id_ed25519` (or `~/.ssh/id_rsa`) key.

## 2. Start the stack

Run from the `local/` directory:

```bash
docker compose --env-file ../.env up --build
```

Services will start in order: `postgres` → `api` → `frontend`

| Service  | URL                        |
|----------|----------------------------|
| Frontend | http://localhost:5173      |
| API      | http://localhost:8080      |
| Postgres | localhost:5432             |

## 3. Stop the stack

```bash
# Stop and remove containers (data is preserved)
docker compose --env-file ../.env down

# Stop and remove containers + delete all data (wipes DB)
docker compose --env-file ../.env down -v
```

## 4. View logs

```bash
docker compose --env-file ../.env logs -f
```
