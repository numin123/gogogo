
# gogogo

Simple Golang project with API and database support.

## Setup
1. Install Go (https://go.dev/dl/)
2. Clone this repository
3. Install dependencies:
   go mod download

## Tools
- Go (https://go.dev/)
- Docker (for database, optional)

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
