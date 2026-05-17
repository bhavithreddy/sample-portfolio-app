# Architecture Guide

Technical architecture, design decisions, and implementation details for the DevOps Portfolio application.

## System Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    End User (Browser)                       │
└────────────────────────┬────────────────────────────────────┘
                         │ HTTP/HTTPS
                         │
        ┌────────────────▼────────────────┐
        │    DNS / CDN (Optional)         │
        │   (CloudFront / Akamai)         │
        └────────────────┬────────────────┘
                         │
        ┌────────────────▼────────────────┐
        │   NGINX Ingress Controller      │
        │  (TLS Termination, Routing)     │
        └────────────────┬────────────────┘
                         │
        ┌────────────────▼────────────────┐
        │  Kubernetes Service (ClusterIP) │
        │    (Load Balancing)             │
        └────────────────┬────────────────┘
                         │
    ┌────────────────────┼────────────────────┐
    │                    │                    │
┌───▼────┐          ┌───▼────┐          ┌───▼────┐
│ Pod 1  │          │ Pod 2  │   ...    │ Pod N  │
│  Go    │          │  Go    │          │  Go    │
│ App    │          │ App    │          │ App    │
└────────┘          └────────┘          └────────┘
    │                   │                   │
    └───────────────────┼───────────────────┘
                        │
            ┌───────────▼──────────┐
            │     Metrics &        │
            │  Monitoring Stack    │
            │ (Prometheus/Grafana) │
            └──────────────────────┘
```

## Application Architecture

### Web Application (Go)

**Framework**: Go standard library (`net/http`)

**Components**:
- **Handler Functions**: Route HTTP requests to appropriate handlers
- **Template Rendering**: Serve static HTML files from `templates/` directory
- **Static Assets**: Serve CSS, images, and files from `static/` directory
- **Health Endpoint**: Kubernetes-compatible health check at `/health`

**Code Structure**:
```
main.go
├── pageHandler()      - Generic page serving handler
├── healthCheckHandler() - Kubernetes probe handler
├── getPort()          - Environment-based port config
└── main()             - Application entry point
```

### Frontend

**Technology**: HTML5, CSS3, Vanilla JavaScript

**Features**:
- Responsive grid layouts with CSS Grid
- CSS Variables for theming
- Flexbox for component layout
- Modern animations with CSS transitions
- No external JavaScript dependencies
- Performance-optimized CSS

**Pages**:
1. **Home** - Hero section, featured projects, skills overview
2. **About** - Bio, career timeline, philosophy, certifications
3. **Skills** - Detailed skill breakdown with proficiency levels
4. **Projects** - Project portfolio with technical details
5. **Contact** - Multiple contact methods and availability

## Container Architecture

### Dockerfile Strategy

**Multi-stage Build**:
```dockerfile
# Stage 1: Builder
FROM golang:1.22-alpine AS builder
# Compile application

# Stage 2: Runtime
FROM alpine:latest
# Minimal image with only the binary
```

**Advantages**:
- Reduced final image size (~30MB vs ~300MB)
- Minimal attack surface
- Faster deployment
- No build tools in production

### Docker Compose

**Services**:
- **Portfolio App**: Main Go application
- **Prometheus** (optional): Metrics collection
- **Grafana** (optional): Visualization

**Network**: Isolated bridge network for service communication

## Kubernetes Architecture

### Deployment Strategy

**Replicas**: 2-5 based on HorizontalPodAutoscaler

**Rolling Update**:
- `maxSurge: 1` - Allow 1 extra pod during update
- `maxUnavailable: 0` - Ensure zero downtime

**Pod Anti-Affinity**:
- Spreads pods across different nodes
- Improves reliability and performance

### Resource Management

**Requests**:
```yaml
cpu: 100m      # Minimum guaranteed CPU
memory: 128Mi   # Minimum guaranteed memory
```

**Limits**:
```yaml
cpu: 500m      # Maximum CPU allowed
memory: 512Mi   # Maximum memory allowed
```

### Probes

1. **Liveness Probe**
   - Restarts unhealthy pods
   - HTTP GET `/health` every 30s
   - Failure threshold: 3 consecutive failures

2. **Readiness Probe**
   - Routes traffic only to ready pods
   - HTTP GET `/health` every 10s
   - Failure threshold: 2 consecutive failures

3. **Startup Probe**
   - Gives app time to start
   - HTTP GET `/health` every 5s
   - Up to 30s to become ready

### Scaling

**HorizontalPodAutoscaler**:
- CPU target: 70% utilization
- Memory target: 80% utilization
- Min replicas: 2
- Max replicas: 5
- Scale-up: 100% increase per 30s
- Scale-down: 50% decrease per 60s

### Network Policies

**Ingress**:
- Only from NGINX Ingress Controller
- Port 8080 (HTTP)

**Egress**:
- DNS queries (port 53/UDP)
- HTTPS connections (port 443/TCP)

### RBAC

**ServiceAccount**: `portfolio`

**Role**: Read-only access to ConfigMaps in portfolio namespace

**Purpose**: Minimal permissions for least privilege

## Data Flow

### Request Processing

```
1. Client Browser
   ↓
2. HTTPS Request (via Ingress)
   ↓
3. NGINX TLS Termination
   ↓
4. Service Load Balancer
   ↓
5. Pod Selection (Round-robin)
   ↓
6. Go Handler Function
   ↓
7. Template/Static File Serving
   ↓
8. Response to Client
```

### Health Check Flow

```
1. Kubernetes Probe Request
   ↓
2. HTTP GET /health
   ↓
3. healthCheckHandler()
   ↓
4. JSON Response {"status": "healthy"}
   ↓
5. Kubernetes Decision (ready/not ready)
```

## Configuration Management

### Environment Variables

**Application Level**:
- `PORT` - Application port (default: 8080)
- `TZ` - Timezone (default: UTC)

**Personal Information** (via .env):
- `YOUR_NAME`
- `YOUR_EMAIL`
- `YOUR_GITHUB`
- `YOUR_LINKEDIN`

**Kubernetes**:
- Injected via ConfigMap
- Can be updated without rebuild
- Pod restart required for changes

## Security Architecture

### Defense in Depth

1. **Container Level**
   - Non-root user (uid: 1000)
   - Read-only root filesystem
   - No privileged access
   - Minimal base image (Alpine)
   - No shell access

2. **Kubernetes Level**
   - Network policies
   - RBAC restrictions
   - Pod security policies
   - Resource quotas

3. **Application Level**
   - Input validation
   - Secure headers
   - Proper error handling
   - Logging and monitoring

4. **Network Level**
   - TLS/HTTPS encryption
   - Ingress security
   - Network segmentation
   - DDoS protection

### Secret Management

**Not Implemented** in this basic version.

**Recommended for Production**:
- Kubernetes Secrets
- HashiCorp Vault
- AWS Secrets Manager
- Azure Key Vault

## Monitoring & Observability

### Health Checks

**Endpoint**: `GET /health`

**Response**:
```json
{"status": "healthy"}
```

**Use Cases**:
- Kubernetes liveness/readiness probes
- Load balancer health checks
- Uptime monitoring

### Logging

**Docker**:
```bash
docker logs -f portfolio
```

**Kubernetes**:
```bash
kubectl logs -f pod/portfolio
kubectl logs -f deploy/portfolio
```

**Log Drivers** (configurable):
- stdout/stderr
- File logging
- CloudWatch / Azure Monitor
- Splunk / ELK Stack

### Metrics (Optional)

**Prometheus Integration**:
- Enable in `docker-compose.yml`
- Metrics endpoint: `http://localhost:8080/metrics`
- Scrape interval: 15s

**Metrics Available**:
- HTTP request duration
- Request count by endpoint
- Response status codes
- Go runtime metrics

### Alerting (Recommended)

**Prometheus AlertRules**:
```yaml
- alert: HighErrorRate
  expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.05

- alert: PodNotReady
  expr: kube_pod_status_ready{namespace="portfolio"} == 0
```

## Performance Optimization

### Frontend Performance

1. **CSS Optimization**
   - CSS Variables for efficient theming
   - Minimal CSS (~15KB)
   - No external dependencies
   - Single stylesheet

2. **No JavaScript**
   - Fully functional without JS
   - Better core web vitals
   - Faster page load

3. **Asset Size**
   - Efficient HTML structure
   - Optimized CSS with compression
   - Minimal DOM tree

### Backend Performance

1. **Go Runtime**
   - Efficient HTTP handling
   - Minimal memory footprint
   - Fast file serving
   - Concurrent request handling

2. **Resource Limits**
   - CPU: 500m max
   - Memory: 512Mi max
   - Prevents runaway processes

3. **Caching**
   - Browser caching (via headers)
   - CDN caching (optional)
   - Static file optimization

## Disaster Recovery

### Backup Strategy

**Code**: Version controlled in Git
**Configuration**: In ConfigMaps
**Data**: None (stateless application)

### High Availability

- **Multi-replica deployment**: 2-5 pods
- **Pod anti-affinity**: Spread across nodes
- **Auto-scaling**: Handles load spikes
- **Rolling updates**: Zero-downtime deployments
- **Health probes**: Automatic pod replacement

### Recovery Procedures

**Pod Failure**:
- Automatic replacement by Kubernetes
- No manual intervention needed
- Health probes ensure readiness

**Node Failure**:
- Pods rescheduled to healthy nodes
- Anti-affinity ensures distribution
- No data loss (stateless)

**Cluster Failure**:
- Multi-cluster deployment (future)
- Traffic failover via DNS
- RTO: ~5 minutes

## Scalability

### Horizontal Scaling

**Automatic** (HPA):
- Scale up: CPU > 70% or Memory > 80%
- Scale down: CPU < 50% and Memory < 60%

**Manual**:
```bash
kubectl scale deployment/portfolio --replicas=10
```

### Vertical Scaling

**Increase resource limits**:
```yaml
resources:
  limits:
    cpu: 1000m
    memory: 1Gi
```

### Load Testing

**Simulate load**:
```bash
ab -n 1000 -c 10 http://portfolio.example.com/
wrk -t4 -c100 -d30s http://portfolio.example.com/
```

## Best Practices Implemented

✅ Container security (non-root, minimal image)
✅ Kubernetes best practices (health probes, resource limits)
✅ RBAC and least privilege
✅ Multi-stage Docker builds
✅ Environment-based configuration
✅ Health check endpoints
✅ Responsive design
✅ Performance optimized
✅ Git version control
✅ Infrastructure as code
✅ CI/CD pipeline
✅ Code quality checks
✅ Security scanning

## Future Enhancements

- [ ] Database integration (optional)
- [ ] Caching layer (Redis)
- [ ] Distributed tracing (Jaeger)
- [ ] Advanced monitoring (Prometheus metrics)
- [ ] Multi-cluster deployment
- [ ] GitOps integration (ArgoCD)
- [ ] Helm charts
- [ ] API endpoints
- [ ] Authentication/Authorization
- [ ] Dynamic content management

---

For deployment details, see [DEPLOYMENT.md](DEPLOYMENT.md)