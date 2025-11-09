# Gophiway 🛒⚡

A modern, secure, and high-performance full-stack e-commerce framework built with Go and React.

## 🚀 Features

- **Security First**: Built-in authentication, authorization, and security best practices
- **High Performance**: Go backend with efficient concurrent processing
- **Modern Frontend**: React 18 + TypeScript + Tailwind CSS 4
- **Type Safe**: End-to-end type safety with TypeScript and Go's static typing
- **Developer Friendly**: Hot reload, clean architecture, comprehensive documentation
- **Production Ready**: Docker support, CI/CD ready, scalable architecture

## 🏗️ Architecture

### Backend (Go)

- **Framework**: Fiber (Express-inspired web framework)
- **Database**: PostgreSQL with GORM
- **Cache**: Redis
- **Authentication**: JWT with refresh tokens
- **Storage**: MinIO (S3-compatible)
- **Architecture**: Clean Architecture with dependency injection

### Frontend (React + TypeScript)

- **Build Tool**: Vite
- **Styling**: Tailwind CSS 4
- **State Management**: Zustand + TanStack Query
- **Routing**: React Router v6
- **Forms**: React Hook Form + Zod
- **HTTP Client**: Axios with interceptors

## 📋 Prerequisites

- **Go** 1.21 or higher
- **Node.js** 18 or higher
- **Docker** and **Docker Compose** (for local development)
- **PostgreSQL** 15+ (if not using Docker)
- **Redis** 7+ (if not using Docker)

## 🛠️ Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/Shihasz/gophiway.git
cd gophiway
```

### 2. Start infrastructure services

```bash
docker-compose up -d
```

### 3. Set up backend

```bash
cd backend
cp .env.example .env
# Edit .env with your configuration
go mod download
go run cmd/api/main.go
```

Backend will be available at `http://localhost:8080`

### 4. Set up frontend

```bash
cd frontend
cp .env.example .env
# Edit .env with your configuration
npm install
npm run dev
```

Frontend will be available at `http://localhost:5173`

## 🧪 Testing

### Backend Tests

```bash
cd backend
go test ./...
```

### Frontend Tests

```bash
cd frontend
npm run test
```

## 📦 Building for Production

### Backend

```bash
cd backend
go build -o bin/api cmd/api/main.go
```

### Frontend

```bash
cd frontend
npm run build
```

## 🔧 Available Make Commands

```bash
make help          # Show all available commands
make dev           # Start development environment
make test          # Run all tests
make build         # Build both backend and frontend
make docker-up     # Start Docker services
make docker-down   # Stop Docker services
make migrate-up    # Run database migrations
make migrate-down  # Rollback database migrations
```

## 📚 Documentation

- [Backend Documentation](./backend/README.md)
- [Frontend Documentation](./frontend/README.md)
- [API Documentation](./docs/api.md)
- [Architecture Guide](./docs/architecture.md)
- [Security Guide](./docs/security.md)

## 🗂️ Project Structure

```
gophiway/
├── backend/            # Go backend application
│   ├── cmd/            # Application entrypoints
│   ├── internal/       # Private application code
│   ├── pkg/            # Public reusable packages
│   └── migrations/     # Database migrations
├── frontend/           # React frontend application
│   ├── src/            # Source code
│   └── public/         # Static assets
├── docs/               # Documentation
├── scripts/            # Build and deployment scripts
└── docker-compose.yml  # Local development setup
```

## 🔐 Security

Gophiway implements security best practices:

- JWT authentication with refresh tokens
- Password hashing with bcrypt
- Rate limiting and request throttling
- CORS protection
- XSS and CSRF protection
- SQL injection prevention
- Secure headers (CSP, HSTS, etc.)
- Input validation and sanitization
- Secure payment integration practices (designed to support PCI DSS compliance via providers like Stripe, Braintree, etc.)

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](./CONTRIBUTING.md) for details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## 🙏 Acknowledgments

Built with ❤️ using:

- [Go](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [React](https://react.dev/)
- [Vite](https://vitejs.dev/)
- [Tailwind CSS](https://tailwindcss.com/)

## 📞 Support

- Documentation: [docs/](./docs/)
- Issues: [GitHub Issues](https://github.com/Shihasz/gophiway/issues)
- Discussions: [GitHub Discussions](https://github.com/Shihasz/gophiway/discussions)

---

**Gophiway** - The fast lane for modern e-commerce 🛣️✨
