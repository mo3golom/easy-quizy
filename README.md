e # Easy Quizy

A full-stack quiz application built with Go backend and SvelteKit frontend, featuring unified Docker deployment and flexible development workflows.

## 🚀 Quick Start

### Prerequisites

- **Docker** (for production deployment)
- **Go 1.24+** (for development)
- **Node.js 24.4+** (for development)
- **npm** (for frontend dependencies)

### Production Deployment with Docker
The easiest way to run Easy Quizy is using Docker, which packages both the backend and frontend into a single container.

#### Build and Run

```bash
# Build the Docker image
docker build -t easy-quizy .

# Run the container
docker run -p 3000:3000 easy-quizy
```

The application will be available at `http://localhost:3000`

#### Using Docker Compose (Recommended)

For a complete setup with database and other services:

```bash
cd docker-easy-quizy
docker-compose up -d
```

This starts:
- Easy Quizy application on port 3000
- PostgreSQL database
- MinIO object storage

### Development Setup

For local development, you can run services individually or together using the development script.

#### Quick Development Start

```bash
# Make the development script executable
chmod +x dev.sh

# Start both backend and frontend (default mode)
./dev.sh

# Or explicitly specify full mode
./dev.sh full
```

This will start:
- Go backend on `http://localhost:8080`
- SvelteKit frontend on `http://localhost:5173`

#### Frontend-Only Development

If you have an external API server or want to work only on the frontend:

```bash
# Start only the frontend
./dev.sh client
```

## 📖 Development Script Usage

The `dev.sh` script provides flexible development modes:

### Available Modes

| Mode | Description | Services Started |
|------|-------------|------------------|
| `full` (default) | Full development environment | Go backend + SvelteKit frontend |
| `client` | Frontend-only development | SvelteKit frontend only |

### Examples

```bash
# Start both services (default)
./dev.sh

# Start both services (explicit)
./dev.sh full

# Start only frontend
./dev.sh client

# Show help
./dev.sh --help
```

### What the Script Does

- **Dependency Checking**: Automatically checks for Go, Node.js, and npm
- **Port Validation**: Ensures required ports (8080, 5173) are available
- **Auto-Installation**: Installs missing Go modules and npm packages
- **Environment Setup**: Creates default `.env` files if missing
- **Process Management**: Handles graceful shutdown with Ctrl+C
- **Status Monitoring**: Provides colored output and status updates

## 🏗️ Architecture

### Production Architecture

```
┌─────────────────────────────────────┐
│           Docker Container          │
│  ┌─────────────────────────────────┐ │
│  │     SvelteKit SSR Server        │ │
│  │         (Port 3000)             │ │
│  │  ┌─────────────────────────────┐ │ │
│  │  │      Static Assets          │ │ │
│  │  └─────────────────────────────┘ │ │
│  └─────────────────────────────────┘ │
│  ┌─────────────────────────────────┐ │
│  │       Go Backend Server         │ │
│  │         (Port 8080)             │ │
│  └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

### Development Architecture

```
┌─────────────────────┐    ┌─────────────────────┐
│  Vite Dev Server    │    │   Go Server         │
│    (Port 5173)      │───▶│   (Port 8080)       │
│                     │    │                     │
│  Hot Reload         │    │  Live Reload        │
│  Proxy /api/* ────────────▶ API Endpoints      │
└─────────────────────┘    └─────────────────────┘
```

## 🐳 Docker Deployment

### Build Options

#### Basic Build
```bash
docker build -t easy-quizy .
```

#### Build with Custom Tag
```bash
docker build -t easy-quizy:v1.0.0 .
```

#### Build with Build Args
```bash
docker build \
  --build-arg PUBLIC_API_BASE_URL="" \
  -t easy-quizy .
```

### Run Options

#### Basic Run
```bash
docker run -p 3000:3000 easy-quizy
```

#### Run with Environment Variables
```bash
docker run \
  -p 3000:3000 \
  -e SERVER_PORT=8080 \
  -e DB_HOST=your-db-host \
  -e DB_USER=your-db-user \
  -e DB_PASSWORD=your-db-password \
  easy-quizy
```

#### Run with Volume Mounts
```bash
docker run \
  -p 3000:3000 \
  -v $(pwd)/.env:/app/.env \
  easy-quizy
```

#### Run in Background (Detached)
```bash
docker run -d -p 3000:3000 --name easy-quizy-app easy-quizy
```

### Docker Compose Examples

#### Basic Setup
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - DB_HOST=postgres
      - DB_USER=postgres
      - DB_PASSWORD=postgres
```

#### With Database
```yaml
version: '3.8'
services:
  app:
    build: .
    ports:
      - "3000:3000"
    depends_on:
      - postgres
    environment:
      - DB_HOST=postgres
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=quizy
  
  postgres:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=quizy
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

## ⚙️ Configuration

### Environment Variables

#### Backend Configuration (.env)
```bash
# Server Configuration
SERVER_PORT=8080

# Database Configuration
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_HOST="localhost"
DB_PORT="5432"
DB_NAME="postgres"
DB_SSL="disable"
```

#### Frontend Configuration (web/.env)
```bash
# API Base URL
# Leave empty for production (uses relative URLs)
# Set for development to point to Go backend
PUBLIC_API_BASE_URL=http://localhost:8080
```

### Port Configuration

| Service | Development Port | Production Port | Purpose |
|---------|------------------|-----------------|---------|
| Go Backend | 8080 | 8080 (internal) | API endpoints |
| SvelteKit Dev | 5173 | - | Development server |
| SvelteKit SSR | - | 3000 | Production server |

## 🛠️ Development

### Manual Setup

If you prefer not to use the development script:

#### Backend Setup
```bash
# Install Go dependencies
go mod download
go mod tidy

# Create .env file (see Configuration section)
cp .env.example .env

# Run the backend
go run cmd/service/main.go
```

#### Frontend Setup
```bash
# Navigate to web directory
cd web

# Install dependencies
npm install

# Create .env file
echo "PUBLIC_API_BASE_URL=http://localhost:8080" > .env

# Run the development server
npm run dev
```

### Building for Production

#### Frontend Build
```bash
cd web
npm run build
```

#### Backend Build
```bash
go build -o go-server ./cmd/service/main.go
```

### Testing

#### Backend Tests
```bash
go test ./...
```

#### Frontend Tests
```bash
cd web
npm run test
```

## 🔧 Troubleshooting

### Common Issues

#### Port Already in Use

**Problem**: Error message about port 8080 or 5173 being in use.

**Solutions**:
```bash
# Find process using the port
lsof -i :8080
lsof -i :5173

# Kill the process (replace PID with actual process ID)
kill -9 <PID>

# Or kill all processes on a port (macOS/Linux)
sudo lsof -t -i:8080 | xargs kill -9
```

#### Docker Build Fails

**Problem**: Docker build fails during frontend or backend build stage.

**Solutions**:
```bash
# Clear Docker build cache
docker builder prune

# Build with no cache
docker build --no-cache -t easy-quizy .

# Check Docker logs for specific errors
docker build -t easy-quizy . 2>&1 | tee build.log
```

#### Frontend Can't Connect to Backend

**Problem**: API calls fail with connection errors.

**Solutions**:

1. **Check backend is running**:
   ```bash
   curl http://localhost:8080/api/health
   ```

2. **Verify environment variables**:
   ```bash
   # Check web/.env file
   cat web/.env
   
   # Should contain:
   PUBLIC_API_BASE_URL=http://localhost:8080
   ```

3. **Check CORS configuration** in Go backend
4. **Verify proxy configuration** in `vite.config.ts`

#### Database Connection Issues

**Problem**: Backend fails to connect to database.

**Solutions**:

1. **Check database is running**:
   ```bash
   # For Docker Compose setup
   docker-compose ps
   
   # For local PostgreSQL
   pg_isready -h localhost -p 5432
   ```

2. **Verify database credentials** in `.env` file
3. **Check database exists**:
   ```bash
   psql -h localhost -U postgres -l
   ```

#### Permission Denied Errors

**Problem**: Permission denied when running scripts or Docker commands.

**Solutions**:
```bash
# Make scripts executable
chmod +x dev.sh
chmod +x start.sh

# Fix Docker permissions (Linux)
sudo usermod -aG docker $USER
# Then logout and login again
```

#### Node.js/npm Issues

**Problem**: npm install fails or Node.js version conflicts.

**Solutions**:
```bash
# Clear npm cache
npm cache clean --force

# Delete node_modules and reinstall
rm -rf web/node_modules
cd web && npm install

# Use specific Node.js version with nvm
nvm use 18
nvm use 20
```

#### Go Module Issues

**Problem**: Go dependencies fail to download or build.

**Solutions**:
```bash
# Clean Go module cache
go clean -modcache

# Re-download dependencies
go mod download
go mod tidy

# Verify Go version
go version  # Should be 1.21 or later
```

### Docker-Specific Troubleshooting

#### Container Exits Immediately

**Problem**: Docker container starts but exits immediately.

**Solutions**:
```bash
# Check container logs
docker logs <container-id>

# Run container interactively for debugging
docker run -it easy-quizy sh

# Check if both services start properly
docker run -p 3000:3000 easy-quizy
```

#### Services Not Accessible

**Problem**: Can't access the application on port 3000.

**Solutions**:
```bash
# Check if container is running
docker ps

# Check port mapping
docker port <container-id>

# Test internal connectivity
docker exec -it <container-id> wget -qO- http://localhost:3000
```

#### Build Performance Issues

**Problem**: Docker build takes too long.

**Solutions**:
```bash
# Use BuildKit for faster builds
DOCKER_BUILDKIT=1 docker build -t easy-quizy .

# Use multi-core builds
docker build --build-arg BUILDKIT_INLINE_CACHE=1 -t easy-quizy .
```

### Getting Help

If you encounter issues not covered here:

1. **Check the logs**: Both development script and Docker provide detailed logging
2. **Verify prerequisites**: Ensure all required tools are installed and up-to-date
3. **Check port availability**: Make sure required ports are not in use
4. **Review configuration**: Verify all environment variables are set correctly
5. **Test components individually**: Try running backend and frontend separately

### Debug Mode

For additional debugging information:

```bash
# Enable verbose logging in development script
DEBUG=1 ./dev.sh

# Run Docker with debug output
docker run -e DEBUG=1 -p 3000:3000 easy-quizy
```

## 📁 Project Structure

```
easy-quizy/
├── api/                    # API route handlers
├── cmd/service/           # Go application entry point
├── internal/              # Internal Go packages
│   ├── contracts/         # Interface definitions
│   ├── model/            # Data models
│   ├── repositories/     # Data access layer
│   └── usecase/          # Business logic
├── migrations/           # Database migrations
├── pkg/                  # Shared Go packages
├── web/                  # SvelteKit frontend
│   ├── src/              # Source code
│   ├── static/           # Static assets
│   └── build/            # Production build output
├── docker-easy-quizy/    # Docker Compose setup
├── .env                  # Backend environment variables
├── dev.sh               # Development script
├── Dockerfile           # Multi-stage Docker build
├── start.sh            # Production startup script
└── README.md           # This file
```