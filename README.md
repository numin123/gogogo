
# gogogo

Simple Golang project with API and database support.

## Setup
1. Install Go (https://go.dev/dl/)
2. Clone this repository
3. Install dependencies:
   go mod download



## Example .env
Create a `.env` file in the project root with:
```
JWT_SECRET_KEY=your_secret_key
DATABASE_URL=postgres://gogogo_user:gogogo_pass@localhost:5432/gogogo_db?sslmode=disable
SERVER_PORT=8080
```

## Dependencies
This project uses Go modules to manage dependencies. Main modules used:

- `github.com/gin-gonic/gin`: Web API framework
- `gorm.io/gorm`: ORM for Go
- `gorm.io/driver/postgres`: PostgreSQL support for GORM
- `github.com/golang-jwt/jwt/v5`: JWT authentication
- `github.com/joho/godotenv`: Load .env files

## How to Start
Run the application:
```
go run ./cmd/main.go
```

## Using Docker
This project includes a `docker-compose.yml` for running a local PostgreSQL database and managing migrations.

### Start Database
```
docker compose up db
```

### Run Migrations
```
docker compose run --rm migrate
```

### Rollback Last Migration
```
docker compose run --rm rollback
```

### Seed Database
```
docker compose run --rm seed
```

## Development
- Source code: `cmd/`, `internal/`
- Database migrations: `migrations/`
- SQL seed: `seed.sql`
- Docker Compose: `docker-compose.yml`

## Other
- Dependencies managed by `go.mod` and `go.sum`
- Configure database connection in your code or with environment variables
- Default DB credentials (see `docker-compose.yml`):
  - user: gogogo_user
  - password: gogogo_pass
  - db: gogogo_db
