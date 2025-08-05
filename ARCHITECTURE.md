# Easy-Quizy Architecture

## Overview
Full-stack quiz application for Telegram Mini Apps with Go backend and SvelteKit frontend.

**Stack**: Go + Gin + PostgreSQL + SvelteKit + Telegram SDK

## Architecture

```
┌─────────────────┐    HTTP/REST    ┌─────────────────┐
│   SvelteKit     │◄──────────────►│   Go Backend    │
│   Frontend      │                 │   (Gin + SQLX)  │
│                 │                 │                 │
│ • Components    │                 │ • API Layer     │
│ • Stores        │                 │ • Business      │
│ • Telegram SDK  │                 │ • Repository    │
└─────────────────┘                 └─────────────────┘
                                             │
                                             ▼
                                    ┌─────────────────┐
                                    │   PostgreSQL    │
                                    │   Database      │
                                    └─────────────────┘
```

## Core Components

### Backend (Go)
- **API Layer** (`api/v1/game`): HTTP handlers, request/response DTOs
- **Business Layer** (`internal/usecase`): Game logic, user management
- **Repository Layer** (`internal/repositories`): Database operations
- **Models** (`internal/model`): Core entities (Game, User, GameSession)

### Frontend (SvelteKit)
- **Pages** (`src/routes`): File-based routing
- **Components** (`src/lib/components`): Reusable UI (QuizQuestion, QuizResult)
- **Stores** (`src/lib/stores`): State management
- **API Client** (`src/lib/api`): Backend communication

## Data Model

```
Game (quiz definition)
├── Questions[]
└── ScoreResults[]

User (player identity)
├── Internal UUID
└── External ID (Telegram)

GameSession (play instance)
├── GameID
├── PlayerID
└── Answers[]
```

## Authentication
Custom header-based auth for Telegram Mini Apps:
- `X-Player-ID`: External user identifier
- `X-Source`: Platform source (telegram)
- Middleware converts to internal User UUID

## Deployment

### Development
```bash
./dev.sh          # Both services
./dev.sh client   # Frontend only
```

### Production
```bash
docker build -t easy-quizy .
docker run -p 3000:3000 easy-quizy
```

## Key Design Decisions

1. **Telegram-First**: Built specifically for Telegram Mini Apps
2. **Layered Backend**: Clean separation of concerns (API → Business → Data)
3. **JSON Quiz Data**: Questions stored in `quizes/*.json` files
4. **Proxy Setup**: SvelteKit proxies `/api/*` to Go backend
5. **Transaction Management**: Uses `go-transaction-manager` for data consistency

## Development Notes

- Quiz data loaded from JSON files at startup
- Frontend uses Telegram SDK for native features (haptics, theme)
- CORS configured for both development and production
- Database migrations in `migrations/` directory