.PHONY: help dev test build clean docker-up docker-down migrate-up migrate-down backend frontend

# Default target
help:
	@echo "Gophiway - Available Commands:"
	@echo "  make dev           - Start development environment"
	@echo "  make test          - Run all tests"
	@echo "  make build         - Build backend and frontend"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-up     - Start Docker services"
	@echo "  make docker-down   - Stop Docker services"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback database migrations"
	@echo "  make backend       - Run backend server"
	@echo "  make frontend      - Run frontend dev server"
	@echo "  make lint          - Run linters"

# Start development environment
dev: docker-up
	@echo "Starting development environment..."
	@trap 'make docker-down' EXIT; \
	(cd backend && go run cmd/api/main.go) & \
	(cd frontend && npm run dev)

# Run all tests
test:
	@echo "Running backend tests..."
	cd backend && go test ./... -v -cover
	@echo "Running frontend tests..."
	cd frontend && npm run test

# Build both backend and frontend
build: build-backend build-frontend

build-backend:
	@echo "Building backend..."
	cd backend && go build -o bin/api cmd/api/main.go

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm run build

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf backend/bin
	rm -rf frontend/dist
	rm -rf frontend/build

# Docker commands
docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d
	@echo "Waiting for services to be ready..."
	sleep 5

docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

docker-clean:
	@echo "Cleaning Docker volumes..."
	docker-compose down -v

# Database migrations
migrate-up:
	@echo "Running database migrations..."
	cd backend && go run cmd/migrate/main.go up

migrate-down:
	@echo "Rolling back database migrations..."
	cd backend && go run cmd/migrate/main.go down

# Run backend server
backend:
	@echo "Starting backend server..."
	cd backend && go run cmd/api/main.go

# Run frontend dev server
frontend:
	@echo "Starting frontend dev server..."
	cd frontend && npm run dev

# Linting
lint:
	@echo "Running backend linter..."
	cd backend && golangci-lint run
	@echo "Running frontend linter..."
	cd frontend && npm run lint

# Install dependencies
install:
	@echo "Installing backend dependencies..."
	cd backend && go mod download
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# Generate code
generate:
	@echo "Generating code..."
	cd backend && go generate ./...

# Database seed
seed:
	@echo "Seeding database..."
	cd backend && go run cmd/seed/main.go

# Format code
fmt:
	@echo "Formatting backend code..."
	cd backend && go fmt ./...
	@echo "Formatting frontend code..."
	cd frontend && npm run format
