.PHONY: help build run test clean docker-build docker-run k8s-deploy k8s-delete fmt lint

# Variables
BINARY_NAME=portfolio
DOCKER_IMAGE=YOUR_GITHUB/devops-portfolio
DOCKER_TAG=latest
GO_VERSION=$(shell go version | awk '{print $$3}')

help: ## Display this help screen
	@echo "DevOps Portfolio - Makefile Commands"
	@echo "===================================="
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the Go application
	@echo "🔨 Building application..."
	@go build -o $(BINARY_NAME) -v
	@echo "✅ Build complete: ./$(BINARY_NAME)"

run: build ## Build and run the application
	@echo "🚀 Running application..."
	@./$(BINARY_NAME)

test: ## Run all tests
	@echo "🧪 Running tests..."
	@go test -v -race -coverprofile=coverage.txt ./...
	@go tool cover -func=coverage.txt | tail -1

test-coverage: ## Run tests and display coverage report
	@echo "📊 Running tests with coverage..."
	@go test -v -coverprofile=coverage.txt ./...
	@go tool cover -html=coverage.txt -o coverage.html
	@echo "📈 Coverage report generated: coverage.html"

fmt: ## Format Go code
	@echo "📝 Formatting code..."
	@go fmt ./...
	@echo "✅ Format complete"

lint: ## Run Go linter
	@echo "🔍 Linting code..."
	@go vet ./...
	@echo "✅ Lint complete"

clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.txt coverage.html
	@echo "✅ Cleanup complete"

docker-build: ## Build Docker image
	@echo "🐳 Building Docker image: $(DOCKER_IMAGE):$(DOCKER_TAG)"
	@docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	@echo "✅ Docker image built successfully"

docker-run: docker-build ## Build and run Docker container
	@echo "🐳 Running Docker container..."
	@docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

docker-stop: ## Stop Docker container
	@echo "⛔ Stopping Docker container..."
	@docker stop devops-portfolio 2>/dev/null || true
	@echo "✅ Container stopped"

docker-push: docker-build ## Push Docker image to registry
	@echo "📤 Pushing Docker image to registry..."
	@docker push $(DOCKER_IMAGE):$(DOCKER_TAG)
	@echo "✅ Image pushed successfully"

compose-up: ## Start services with Docker Compose
	@echo "🚀 Starting services with Docker Compose..."
	@docker-compose up -d
	@echo "✅ Services started"
	@echo "📌 Portfolio: http://localhost:8080"

compose-down: ## Stop Docker Compose services
	@echo "⛔ Stopping Docker Compose services..."
	@docker-compose down
	@echo "✅ Services stopped"

compose-logs: ## View Docker Compose logs
	@docker-compose logs -f portfolio

k8s-deploy: ## Deploy to Kubernetes cluster
	@echo "☸️ Deploying to Kubernetes..."
	@kubectl apply -f k8s/
	@echo "✅ Kubernetes deployment complete"
	@echo "📌 Check status: kubectl get pods -n portfolio"

k8s-delete: ## Delete Kubernetes deployment
	@echo "☸️ Deleting Kubernetes deployment..."
	@kubectl delete -f k8s/
	@echo "✅ Kubernetes resources deleted"

k8s-status: ## Check Kubernetes deployment status
	@echo "☸️ Checking Kubernetes status..."
	@kubectl get pods -n portfolio
	@kubectl get svc -n portfolio
	@kubectl get ingress -n portfolio

k8s-logs: ## View Kubernetes pod logs
	@kubectl logs -n portfolio -l app=portfolio -f --tail=50

k8s-port-forward: ## Port forward to Kubernetes service
	@echo "🔌 Port forwarding to Kubernetes service..."
	@echo "📌 Access at: http://localhost:8080"
	@kubectl port-forward -n portfolio svc/portfolio 8080:80

k8s-describe: ## Describe Kubernetes deployment
	@kubectl describe deployment -n portfolio portfolio

install-deps: ## Install Go dependencies
	@echo "📦 Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "✅ Dependencies installed"

version: ## Display Go version and binary info
	@echo "📦 Project Information"
	@echo "====================="
	@echo "Go Version: $(GO_VERSION)"
	@echo "Binary Name: $(BINARY_NAME)"
	@echo "Docker Image: $(DOCKER_IMAGE):$(DOCKER_TAG)"
	@echo ""

all: fmt lint test build ## Run all checks and build
	@echo "✅ All checks passed and application built"

.DEFAULT_GOAL := help