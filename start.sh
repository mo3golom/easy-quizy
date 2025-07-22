#!/bin/sh

# Service startup orchestration script for Docker container
# Runs both Go backend and Node.js SSR server concurrently with proper process management

set -e

# Color codes for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BACKEND_PORT=${SERVER_PORT:-8080}
FRONTEND_PORT=${FRONTEND_PORT:-3000}
SHUTDOWN_TIMEOUT=10

# Process IDs
BACKEND_PID=""
FRONTEND_PID=""

# Cleanup function for graceful shutdown
cleanup() {
    echo -e "${YELLOW}Received shutdown signal, initiating graceful shutdown...${NC}"
    
    if [ -n "$BACKEND_PID" ] && kill -0 "$BACKEND_PID" 2>/dev/null; then
        echo -e "${BLUE}Stopping Go backend server (PID: $BACKEND_PID)...${NC}"
        kill -TERM "$BACKEND_PID" 2>/dev/null || true
    fi
    
    if [ -n "$FRONTEND_PID" ] && kill -0 "$FRONTEND_PID" 2>/dev/null; then
        echo -e "${BLUE}Stopping SvelteKit SSR server (PID: $FRONTEND_PID)...${NC}"
        kill -TERM "$FRONTEND_PID" 2>/dev/null || true
    fi
    
    # Wait for graceful shutdown with timeout
    local timeout=$SHUTDOWN_TIMEOUT
    while [ $timeout -gt 0 ]; do
        local running=0
        
        if [ -n "$BACKEND_PID" ] && kill -0 "$BACKEND_PID" 2>/dev/null; then
            running=$((running + 1))
        fi
        
        if [ -n "$FRONTEND_PID" ] && kill -0 "$FRONTEND_PID" 2>/dev/null; then
            running=$((running + 1))
        fi
        
        if [ $running -eq 0 ]; then
            echo -e "${GREEN}All services stopped gracefully${NC}"
            exit 0
        fi
        
        sleep 1
        timeout=$((timeout - 1))
    done
    
    # Force kill if graceful shutdown failed
    echo -e "${RED}Graceful shutdown timeout, force killing processes...${NC}"
    if [ -n "$BACKEND_PID" ]; then
        kill -KILL "$BACKEND_PID" 2>/dev/null || true
    fi
    if [ -n "$FRONTEND_PID" ]; then
        kill -KILL "$FRONTEND_PID" 2>/dev/null || true
    fi
    
    exit 1
}

# Set up signal handlers
trap cleanup TERM INT QUIT

# Function to check if port is available
check_port() {
    local port=$1
    if netstat -ln 2>/dev/null | grep -q ":$port "; then
        echo -e "${RED}Error: Port $port is already in use${NC}"
        return 1
    fi
    return 0
}

# Function to wait for service to be ready
wait_for_service() {
    local port=$1
    local service_name=$2
    local max_attempts=30
    local attempt=1
    
    echo -e "${BLUE}Waiting for $service_name to be ready on port $port...${NC}"
    
    while [ $attempt -le $max_attempts ]; do
        if netstat -ln 2>/dev/null | grep -q ":$port "; then
            echo -e "${GREEN}$service_name is ready on port $port${NC}"
            return 0
        fi
        
        sleep 1
        attempt=$((attempt + 1))
    done
    
    echo -e "${RED}$service_name failed to start within timeout${NC}"
    return 1
}

# Main startup sequence
echo -e "${GREEN}=== Easy Quizy Docker Container Startup ===${NC}"
echo -e "${BLUE}Backend port: $BACKEND_PORT${NC}"
echo -e "${BLUE}Frontend port: $FRONTEND_PORT${NC}"

# Check if ports are available (in case of restart)
if ! check_port "$BACKEND_PORT"; then
    echo -e "${YELLOW}Warning: Backend port $BACKEND_PORT appears to be in use${NC}"
fi

if ! check_port "$FRONTEND_PORT"; then
    echo -e "${YELLOW}Warning: Frontend port $FRONTEND_PORT appears to be in use${NC}"
fi

# Start Go backend server
echo -e "${GREEN}Starting Go backend server on port $BACKEND_PORT...${NC}"
export SERVER_PORT="$BACKEND_PORT"
./go-server &
BACKEND_PID=$!

if ! kill -0 "$BACKEND_PID" 2>/dev/null; then
    echo -e "${RED}Failed to start Go backend server${NC}"
    exit 1
fi

echo -e "${BLUE}Go backend server started with PID: $BACKEND_PID${NC}"

# Wait for backend to be ready
if ! wait_for_service "$BACKEND_PORT" "Go backend"; then
    echo -e "${RED}Backend service failed to start properly${NC}"
    cleanup
    exit 1
fi

# Start SvelteKit SSR server
echo -e "${GREEN}Starting SvelteKit SSR server on port $FRONTEND_PORT...${NC}"
cd web
export PORT="$FRONTEND_PORT"
node build/index.js &
FRONTEND_PID=$!

if ! kill -0 "$FRONTEND_PID" 2>/dev/null; then
    echo -e "${RED}Failed to start SvelteKit SSR server${NC}"
    cleanup
    exit 1
fi

echo -e "${BLUE}SvelteKit SSR server started with PID: $FRONTEND_PID${NC}"

# Wait for frontend to be ready
if ! wait_for_service "$FRONTEND_PORT" "SvelteKit SSR"; then
    echo -e "${RED}Frontend service failed to start properly${NC}"
    cleanup
    exit 1
fi

echo -e "${GREEN}=== All services started successfully ===${NC}"
echo -e "${GREEN}Backend: http://localhost:$BACKEND_PORT${NC}"
echo -e "${GREEN}Frontend: http://localhost:$FRONTEND_PORT${NC}"
echo -e "${BLUE}Container is ready to serve requests${NC}"

# Monitor both processes
while true; do
    # Check if backend is still running
    if [ -n "$BACKEND_PID" ] && ! kill -0 "$BACKEND_PID" 2>/dev/null; then
        echo -e "${RED}Go backend server (PID: $BACKEND_PID) has stopped unexpectedly${NC}"
        cleanup
        exit 1
    fi
    
    # Check if frontend is still running
    if [ -n "$FRONTEND_PID" ] && ! kill -0 "$FRONTEND_PID" 2>/dev/null; then
        echo -e "${RED}SvelteKit SSR server (PID: $FRONTEND_PID) has stopped unexpectedly${NC}"
        cleanup
        exit 1
    fi
    
    # Sleep for a short interval before checking again
    sleep 5
done