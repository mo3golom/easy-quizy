FROM node:24.4-alpine AS frontend-builder

ENV PUBLIC_API_BASE_URL=""
WORKDIR /app/web

COPY web/package*.json ./
COPY web/ ./

RUN npm ci --only=production=false
RUN npm run build

# -------------------------------------------
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY ./ ./

RUN go mod download
RUN go mod verify
RUN go build -o go-server ./cmd/service/main.go

# -------------------------------------------
FROM node:24.4-alpine AS runtime

WORKDIR /app

# Create non-root user for security
RUN addgroup -g 1001 -S nodejs && \
    adduser -S nextjs -u 1001

COPY --from=frontend-builder --chown=nextjs:nodejs /app/web/build ./web/build
COPY --from=frontend-builder --chown=nextjs:nodejs /app/web/package*.json ./web/
COPY --from=backend-builder --chown=nextjs:nodejs /app/go-server ./go-server


WORKDIR /app/web
RUN npm ci --only=production && npm cache clean --force

# Copy and configure startup script
WORKDIR /app
COPY --chown=nextjs:nodejs start.sh ./start.sh
RUN chmod +x start.sh

# Switch to non-root user
USER nextjs

# Expose port 3000 for external access (SSR server)
EXPOSE 3000

# Start both services
CMD ["./start.sh"]