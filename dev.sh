#!/bin/bash

# Development script for Easy Quizy
# Supports multiple modes: full (default), client

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Global variables for process management
BACKEND_PID=""
FRONTEND_PID=""
TUNNEL_PID=""

# Function to print colored output
print_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to show usage
show_usage() {
    echo "Usage: $0 [MODE] [OPTIONS]"
    echo ""
    echo "Modes:"
    echo "  full     Start both Go backend and SvelteKit frontend (default)"
    echo "  client   Start only frontend with external API configuration"
    echo ""
    echo "Options:"
    echo "  --tunnel Start localtunnel for external access (requires localtunnel to be installed)"
    echo ""
    echo "Examples:"
    echo "  $0                    # Start in full mode"
    echo "  $0 full               # Start in full mode"
    echo "  $0 client             # Start in client-only mode"
    echo "  $0 --tunnel           # Start in full mode with localtunnel"
    echo "  $0 full --tunnel      # Start in full mode with localtunnel"
    echo "  $0 client --tunnel    # Start in client-only mode with localtunnel"
}

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Function to check if a port is available
is_port_available() {
    local port=$1
    if command_exists lsof; then
        ! lsof -i :$port >/dev/null 2>&1
    elif command_exists netstat; then
        ! netstat -ln | grep -q ":$port "
    else
        # Fallback: try to bind to the port
        (echo >/dev/tcp/localhost/$port) >/dev/null 2>&1 && return 1 || return 0
    fi
}

# Function to check and install localtunnel
check_localtunnel() {
    if ! command_exists lt; then
        print_warning "localtunnel is not installed. Installing globally..."
        npm install -g localtunnel
        if ! command_exists lt; then
            print_error "Failed to install localtunnel. Please install it manually: npm install -g localtunnel"
            exit 1
        fi
        print_success "localtunnel installed successfully"
    else
        print_success "localtunnel is already installed"
    fi
}

# Function to check and install dependencies
check_dependencies() {
    local use_tunnel=${1:-false}
    
    print_info "Checking dependencies..."
    
    # Check Go
    if ! command_exists go; then
        print_error "Go is not installed. Please install Go 1.21 or later."
        print_info "Visit: https://golang.org/doc/install"
        exit 1
    fi
    
    # Check Node.js
    if ! command_exists node; then
        print_error "Node.js is not installed. Please install Node.js 18 or later."
        print_info "Visit: https://nodejs.org/"
        exit 1
    fi
    
    # Check npm
    if ! command_exists npm; then
        print_error "npm is not installed. Please install npm."
        exit 1
    fi
    
    print_success "All required tools are available"
    
    # Check localtunnel if needed
    if [ "$use_tunnel" = "true" ]; then
        check_localtunnel
    fi
    
    # Check Go dependencies
    print_info "Checking Go dependencies..."
    if [ ! -f "go.sum" ] || [ "go.mod" -nt "go.sum" ]; then
        print_info "Installing Go dependencies..."
        go mod download
        go mod tidy
    fi
    
    # Check Node.js dependencies
    print_info "Checking Node.js dependencies..."
    if [ ! -d "web/node_modules" ] || [ "web/package.json" -nt "web/node_modules" ]; then
        print_info "Installing Node.js dependencies..."
        cd web
        npm install
        cd ..
    fi
    
    print_success "Dependencies are up to date"
}

# Function to check port availability
check_ports() {
    local mode=$1
    
    if [ "$mode" = "full" ]; then
        # Check backend port (8080)
        if ! is_port_available 8080; then
            print_error "Port 8080 is already in use. Please stop the service using this port."
            print_info "You can find the process using: lsof -i :8080"
            exit 1
        fi
        
        # Check frontend port (5173)
        if ! is_port_available 5173; then
            print_error "Port 5173 is already in use. Please stop the service using this port."
            print_info "You can find the process using: lsof -i :5173"
            exit 1
        fi
    elif [ "$mode" = "client" ]; then
        # Check only frontend port (5173)
        if ! is_port_available 5173; then
            print_error "Port 5173 is already in use. Please stop the service using this port."
            print_info "You can find the process using: lsof -i :5173"
            exit 1
        fi
    fi
}

# Function to start localtunnel
start_tunnel() {
    local subdomain="easy-quizy"
    
    print_info "Starting localtunnel with subdomain: $subdomain..."
    
    # Start localtunnel in background
    lt --port 5173 --subdomain "$subdomain" &
    TUNNEL_PID=$!
    
    # Wait a moment for tunnel to establish
    sleep 3
    if ! kill -0 $TUNNEL_PID 2>/dev/null; then
        print_error "Failed to start localtunnel"
        exit 1
    fi
    
    local tunnel_url="https://$subdomain.loca.lt"
    print_success "Localtunnel started (PID: $TUNNEL_PID)"
    print_info "Tunnel URL: $tunnel_url"
    
    # Store the URL for other functions to use
    TUNNEL_URL="$tunnel_url"
}

# Function to cleanup processes on exit
cleanup() {
    print_info "Shutting down services..."
    
    if [ -n "$TUNNEL_PID" ]; then
        print_info "Stopping localtunnel (PID: $TUNNEL_PID)..."
        kill $TUNNEL_PID 2>/dev/null || true
        wait $TUNNEL_PID 2>/dev/null || true
    fi
    
    if [ -n "$BACKEND_PID" ]; then
        print_info "Stopping Go backend (PID: $BACKEND_PID)..."
        kill $BACKEND_PID 2>/dev/null || true
        wait $BACKEND_PID 2>/dev/null || true
    fi
    
    if [ -n "$FRONTEND_PID" ]; then
        print_info "Stopping SvelteKit frontend (PID: $FRONTEND_PID)..."
        kill $FRONTEND_PID 2>/dev/null || true
        wait $FRONTEND_PID 2>/dev/null || true
    fi
    
    print_success "All services stopped"
    exit 0
}

# Function to start Go backend
start_backend() {
    print_info "Starting Go backend server on port 8080..."
    
    # Ensure .env file exists
    if [ ! -f ".env" ]; then
        print_error ".env file not found."
    fi
    
    # Start Go server in background
    go run cmd/service/main.go &
    BACKEND_PID=$!
    
    # Wait a moment and check if the process is still running
    sleep 2
    if ! kill -0 $BACKEND_PID 2>/dev/null; then
        print_error "Failed to start Go backend server"
        exit 1
    fi
    
    print_success "Go backend server started (PID: $BACKEND_PID)"
}

# Function to start SvelteKit frontend
start_frontend() {
    local external_api=${1:-false}
    
    print_info "Starting SvelteKit frontend server on port 5173..."
    
    if [ ! -f "web/.env" ]; then
        print_error "web/.env file not found."
    fi
    
    # Start frontend server in background
    cd web
    npm run dev &
    FRONTEND_PID=$!
    cd ..
    
    # Wait a moment and check if the process is still running
    sleep 3
    if ! kill -0 $FRONTEND_PID 2>/dev/null; then
        print_error "Failed to start SvelteKit frontend server"
        exit 1
    fi
    
    print_success "SvelteKit frontend server started (PID: $FRONTEND_PID)"
}

# Function to start full mode (both backend and frontend)
start_full_mode() {
    local use_tunnel=${1:-false}
    
    print_info "Starting in FULL mode (backend + frontend)..."
    
    start_backend
    start_frontend false
    
    if [ "$use_tunnel" = "true" ]; then
        start_tunnel
    fi
    
    echo ""
    print_success "Both services are running!"
    print_info "Frontend: http://localhost:5173"
    print_info "Backend API: http://localhost:8080"
    
    if [ "$use_tunnel" = "true" ]; then
        print_info "External access: $TUNNEL_URL"
    fi
    
    print_info "Press Ctrl+C to stop all services"
    echo ""
    
    # Wait for any process to exit
    wait
}

# Function to start client-only mode
start_client_mode() {
    local use_tunnel=${1:-false}
    
    print_info "Starting in CLIENT mode (frontend only)..."
    
    start_frontend true
    
    if [ "$use_tunnel" = "true" ]; then
        start_tunnel
    fi
    
    echo ""
    print_success "Frontend service is running!"
    print_info "Frontend: http://localhost:5173"
    print_info "API calls will be made to the configured external API"
    print_info "Current API URL: $(grep PUBLIC_API_BASE_URL web/.env | cut -d'=' -f2)"
    
    if [ "$use_tunnel" = "true" ]; then
        print_info "External access: $TUNNEL_URL"
        print_warning "Perfect for testing Telegram Mini Apps!"
    fi
    
    print_info "Press Ctrl+C to stop the service"
    echo ""
    
    # Wait for the frontend process
    wait $FRONTEND_PID
}

# Main script logic
main() {
    # Set up signal handlers for cleanup
    trap cleanup SIGINT SIGTERM EXIT
    
    # Parse command line arguments
    MODE=""
    USE_TUNNEL=false
    
    # Parse all arguments
    for arg in "$@"; do
        case $arg in
            --tunnel)
                USE_TUNNEL=true
                ;;
            full|client)
                MODE=$arg
                ;;
            -h|--help|help)
                show_usage
                exit 0
                ;;
        esac
    done
    
    # Set default mode if not specified
    if [ -z "$MODE" ]; then
        MODE="full"
    fi
    
    case $MODE in
        full)
            check_dependencies "$USE_TUNNEL"
            check_ports "full"
            start_full_mode "$USE_TUNNEL"
            ;;
        client)
            check_dependencies "$USE_TUNNEL"
            check_ports "client"
            start_client_mode "$USE_TUNNEL"
            ;;
        -h|--help|help)
            show_usage
            exit 0
            ;;
        *)
            print_error "Invalid mode: $MODE"
            echo ""
            show_usage
            exit 1
            ;;
    esac
}

# Run main function with all arguments
main "$@"