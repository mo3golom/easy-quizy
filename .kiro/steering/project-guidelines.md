# Project Guidelines and Standards

## Development Principles

### 1. Minimal Code Implementation
- Create the absolute minimum code necessary to solve the problem
- Avoid over-engineering and unnecessary abstractions
- Focus on essential functionality only
- Keep implementations simple and direct

### 2. Test Separation
- **Never write tests in the same task where you implement logic**
- Create dedicated, separate tasks specifically for testing
- Implement functionality first, then create comprehensive tests in a follow-up task
- This ensures clear separation of concerns and focused development

### 3. Confidence and Communication
- If you are not very confident about implementation details, **ask the user before proceeding**
- Clarify requirements when uncertain
- Discuss architectural decisions when multiple approaches are viable
- Better to ask than to implement incorrectly

## Technology Stack and Versions

### Frontend (SvelteKit)
- **SvelteKit**: ^2.5.27
- **Svelte**: ^5.0.0
- **TypeScript**: ^5.5.0
- **Vite**: ^5.4.4
- **TailwindCSS**: ^4.0.0 (via @tailwindcss/cli)
- **DaisyUI**: ^5.0.46
- **Telegram Apps SDK**: ^3.11.4
- **Node Types**: ^20.10.0

### Backend (Go)
- **Go Version**: 1.24.2
- **Gin Framework**: v1.10.1
- **SQLX**: v1.3.5
- **PostgreSQL Driver**: v1.10.9 (lib/pq)
- **UUID**: v1.6.0 (google/uuid)
- **Logrus**: v1.9.3
- **Transaction Manager**: v2.0.0 (avito-tech)
- **CORS**: v1.7.6 (gin-contrib/cors)
- **Sentry**: v0.34.0
- **GoDotEnv**: v1.5.1

## Code Standards

### Go Backend
- Use Gin framework for HTTP routing
- Implement clean architecture with separate layers:
  - `api/` - HTTP handlers
  - `internal/usecase/` - Business logic
  - `internal/repositories/` - Data access
  - `internal/model/` - Domain models
  - `internal/contracts/` - Interfaces
- Use SQLX for database operations
- Implement proper error handling with Sentry integration
- Use structured logging with Logrus

### SvelteKit Frontend
- Use TypeScript for all new code
- Leverage Svelte 5's new features (runes, etc.)
- Use TailwindCSS + DaisyUI for styling
- Implement proper type safety
- Use stores for state management
- Follow SvelteKit's file-based routing

## Database
- PostgreSQL as primary database
- Use migrations for schema changes
- UUID for primary keys where applicable
- Proper indexing and constraints

## Development Workflow
1. Implement minimal viable solution
2. Create separate task for comprehensive testing
3. Ask for clarification when uncertain
4. Focus on essential functionality only
5. Use established patterns from existing codebase