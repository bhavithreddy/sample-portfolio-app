# Deployment Guide

Complete guide for deploying the DevOps Portfolio application across different environments.

## Table of Contents
1. [Local Development](#local-development)
2. [Docker Deployment](#docker-deployment)
3. [Docker Compose](#docker-compose)
4. [Kubernetes Deployment](#kubernetes-deployment)
5. [Cloud Platforms](#cloud-platforms)
6. [CI/CD Pipeline](#cicd-pipeline)
7. [Production Checklist](#production-checklist)
8. [Troubleshooting](#troubleshooting)

## Local Development

### Prerequisites
- Go 1.22+
- Git

### Setup

```bash
# Clone repository
git clone https://github.com/YOUR_GITHUB/devops-portfolio.git
cd devops-portfolio

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Edit .env with your personal information
vim .env
```

### Run Application

```bash
# Run directly
go run main.go

# Build and run
go build -o portfolio
./portfolio

# Using Make
make run
```

Application available at: `http://localhost:8080`

### Testing

```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -v -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt

# Using Make
make test
make test-coverage
```

## Docker Deployment

### Prerequisites
- Docker 20.10+
- Docker Build Kit (optional but recommended)

### Build Image

```bash
# Basic build
docker build -t devops-portfolio:latest .

# With specific tag
docker build -t YOUR_GITHUB/devops-portfolio:v1.0.0 .

# Using Make
make docker-build
```

### Run Container

```bash
# Run with default settings
docker run -p 8080:8080 devops-portfolio:latest

# Run with environment variables
docker run -p 8080:8080 \
  -e PORT=8080 \
  -e TZ=UTC \
  devops-portfolio:latest

# Run in background
docker run -d -p 8080:8080 --name portfolio devops-portfolio:latest

# View logs
docker logs portfolio
docker logs -f portfolio

# Stop container
docker stop portfolio
docker rm portfolio

# Using Make
make docker-run
make docker-stop
```

### Push to Registry

```bash
# Docker Hub
docker tag devops-portfolio:latest YOUR_GITHUB/devops-portfolio:latest
docker login
docker push YOUR_GITHUB/devops-portfolio:latest

# With specific version
docker tag devops-portfolio:latest YOUR_GITHUB/devops-portfolio:v1.0.0
docker push YOUR_GITHUB/devops-portfolio:v1.0.0

# Using Make
make docker-push
```

## Docker Compose

### Setup

```bash
# Create environment file
cp .env.example .env
vim .env

# Start services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Using Make
make compose-up
make compose-down
make compose-logs
```

### Services

- **Portfolio**: Main application on port 8080
- **Prometheus**: Metrics collection (optional, uncomment in docker-compose.yml)
- **Grafana**: Visualization (optional, uncomment in docker-compose.yml)

## Kubernetes Deployment

### Prerequisites

- kubectl configured
- Kubernetes cluster (EKS, AKS, GKE, or local)
- NGINX Ingress Controller (for ingress)
- cert-manager (for TLS, optional)

### Installation Steps

1. **Build and push Docker image**
   ```bash
   docker build -t YOUR_GITHUB/devops-portfolio:latest .
   docker push YOUR_GITHUB/devops-portfolio:latest
   ```

2. **Update deployment manifests**
   ```bash
   # Edit k8s/deployment.yaml
   # Replace YOUR_GITHUB with your Docker username
   sed -i 's/YOUR_GITHUB/your-username/g' k8s/deployment.yaml
   sed -i 's/YOUR_GITHUB/your-username/g' k8s/ingress.yaml
   ```

3. **Deploy to cluster**
   ```bash
   # Deploy all resources
   kubectl apply -f k8s/

   # Or deploy individually
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/ingress.yaml

   # Using Make
   make k8s-deploy
   ```

4. **Verify deployment**
   ```bash
   # Check pods
   kubectl get pods -n portfolio
   kubectl get svc -n portfolio
   kubectl get ingress -n portfolio

   # View pod status
   kubectl describe pod -n portfolio POD_NAME

   # View logs
   kubectl logs -n portfolio -l app=portfolio
   kubectl logs -n portfolio -l app=portfolio -f

   # Using Make
   make k8s-status
   make k8s-logs
   ```

### Access Application

```bash
# Port forward to local
kubectl port-forward -n portfolio svc/portfolio 8080:80

# Access at http://localhost:8080

# Using Make
make k8s-port-forward

# Via Ingress (if configured)
# Access via portfolio.example.com
```

### Update Deployment

```bash
# Update image
kubectl set image deployment/portfolio \
  portfolio=YOUR_GITHUB/devops-portfolio:v1.0.1 \
  -n portfolio --record

# Verify rollout
kubectl rollout status deployment/portfolio -n portfolio

# Rollback if needed
kubectl rollout undo deployment/portfolio -n portfolio
```

### Scale Deployment

```bash
# Manual scaling
kubectl scale deployment/portfolio --replicas=3 -n portfolio

# Check HorizontalPodAutoscaler
kubectl get hpa -n portfolio
kubectl describe hpa portfolio -n portfolio
```

### Delete Deployment

```bash
# Delete all resources
kubectl delete -f k8s/

# Or delete namespace (deletes everything in it)
kubectl delete namespace portfolio

# Using Make
make k8s-delete
```

## Cloud Platforms

### AWS (EKS)

1. **Create EKS cluster**
   ```bash
   eksctl create cluster --name portfolio --region us-east-1
   ```

2. **Install NGINX Ingress**
   ```bash
   helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
   helm install ingress-nginx ingress-nginx/ingress-nginx \
     -n ingress-nginx --create-namespace
   ```

3. **Deploy application**
   ```bash
   kubectl apply -f k8s/
   ```

4. **Get load balancer IP**
   ```bash
   kubectl get svc -n ingress-nginx
   ```

### Azure (AKS)

1. **Create AKS cluster**
   ```bash
   az aks create --resource-group myResourceGroup \
     --name portfolioCluster --node-count 3
   ```

2. **Get credentials**
   ```bash
   az aks get-credentials --resource-group myResourceGroup \
     --name portfolioCluster
   ```

3. **Deploy application**
   ```bash
   kubectl apply -f k8s/
   ```

### GCP (GKE)

1. **Create GKE cluster**
   ```bash
   gcloud container clusters create portfolio \
     --zone us-central1-a --num-nodes 3
   ```

2. **Get credentials**
   ```bash
   gcloud container clusters get-credentials portfolio \
     --zone us-central1-a
   ```

3. **Deploy application**
   ```bash
   kubectl apply -f k8s/
   ```

## CI/CD Pipeline

### GitHub Actions

The `.github/workflows/ci-cd.yml` workflow includes:

1. **Code Quality**
   - Go fmt, vet, tests
   - Coverage reporting

2. **Security**
   - Trivy filesystem scan
   - Container image scan

3. **Build**
   - Multi-stage Docker build
   - Push to Docker Hub

4. **Deploy**
   - Update Kubernetes deployment
   - Verify rollout

### Setup

1. **Add secrets to GitHub**
   ```
   Settings > Secrets and variables > Actions > New repository secret
   ```

   Required secrets:
   - `DOCKER_USERNAME`: Your Docker Hub username
   - `DOCKER_PASSWORD`: Your Docker Hub token
   - `KUBE_CONFIG`: Base64 encoded kubeconfig

2. **Trigger workflow**
   - Push to `main` or `develop` branch
   - Push version tag (`v1.0.0`)

## Production Checklist

- [ ] Update `.env` with production values
- [ ] Replace all `YOUR_*` placeholders
- [ ] Test application locally
- [ ] Run all tests and linters
- [ ] Build and test Docker image
- [ ] Push image to registry
- [ ] Review Kubernetes manifests
- [ ] Set up TLS/SSL certificates
- [ ] Configure DNS
- [ ] Set up monitoring (Prometheus/Grafana)
- [ ] Set up logging (ELK/Loki)
- [ ] Configure backups
- [ ] Document deployment steps
- [ ] Set up incident response plan
- [ ] Enable security scanning
- [ ] Test disaster recovery
- [ ] Load testing
- [ ] Performance optimization
- [ ] Security audit

## Troubleshooting

### Application won't start

```bash
# Check logs
docker logs devops-portfolio

# Check Go syntax
go build -o portfolio

# Check file permissions
ls -la templates/
ls -la static/
```

### Kubernetes pod errors

```bash
# Describe pod
kubectl describe pod -n portfolio POD_NAME

# View logs
kubectl logs -n portfolio POD_NAME

# Check events
kubectl get events -n portfolio

# Check resources
kubectl top pods -n portfolio
kubectl top nodes
```

### Connection refused

```bash
# Check if service is running
kubectl get svc -n portfolio

# Port forward manually
kubectl port-forward -n portfolio svc/portfolio 8080:80

# Check ingress
kubectl describe ingress -n portfolio
```

### Health check failing

```bash
# Test endpoint manually
kubectl exec -it -n portfolio POD_NAME -- \
  wget http://localhost:8080/health -O -

# Check port binding
kubectl exec -it -n portfolio POD_NAME -- netstat -tlnp
```

### Out of memory

```bash
# Check memory usage
kubectl top pods -n portfolio

# Increase limits in deployment.yaml
resources:
  limits:
    memory: 1Gi

# Restart deployment
kubectl rollout restart deployment/portfolio -n portfolio
```

## Monitoring

### Health Check

```bash
# Local
curl http://localhost:8080/health

# Kubernetes
kubectl exec -it -n portfolio POD_NAME -- \
  curl http://localhost:8080/health
```

### Metrics

Enable Prometheus and Grafana in `docker-compose.yml` for detailed metrics.

### Logs

```bash
# Docker
docker logs -f portfolio

# Kubernetes
kubectl logs -n portfolio -l app=portfolio -f

# With timestamp
kubectl logs -n portfolio -l app=portfolio --timestamps=true
```

---

For more information, see [README.md](README.md)