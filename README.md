# DevOps Portfolio

A modern, production-ready personal developer portfolio and DevOps showcase platform built with Go, Docker, and Kubernetes.

![License](https://img.shields.io/badge/License-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)
![Docker](https://img.shields.io/badge/Docker-Ready-green.svg)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-green.svg)

## 🚀 Features

- **Modern Dark Theme**: Sleek UI with blue, purple, and cyan gradients
- **Responsive Design**: Mobile-friendly across all devices
- **Multiple Pages**: Home, About, Skills, Projects, Contact
- **Production Ready**: Security-hardened, containerized, cloud-native
- **CI/CD Pipeline**: GitHub Actions with automated testing and deployment
- **Kubernetes Ready**: Complete K8s manifests with best practices
- **DevOps Showcase**: Highlight your Kubernetes, Docker, AWS, Terraform, and other DevOps expertise
- **Performance Optimized**: Fast load times, optimized assets
- **SEO Friendly**: Proper meta tags and semantic HTML
- **Accessibility**: WCAG compliant markup

## 📋 Tech Stack

### Application
- **Language**: Go 1.22
- **Framework**: Standard library (net/http)
- **Frontend**: HTML5, CSS3, Vanilla JavaScript
- **Styling**: Custom modern CSS with CSS variables

### DevOps & Infrastructure
- **Containerization**: Docker (multi-stage builds)
- **Orchestration**: Kubernetes (EKS/AKS/GKE compatible)
- **CI/CD**: GitHub Actions
- **Infrastructure as Code**: Terraform, Kubernetes YAML
- **Monitoring**: Prometheus, Grafana ready

### Security
- **Container Security**: Non-root user, minimal base image (Alpine)
- **RBAC**: Kubernetes RBAC with ServiceAccount
- **Network Security**: Network policies for pod-to-pod communication
- **TLS**: Automatic certificate management with cert-manager
- **Code Scanning**: Trivy for vulnerability scanning

## 🏗️ Project Structure

```
.
├── main.go                      # Application entry point
├── main_test.go                 # Go tests
├── go.mod                       # Go module definition
├── Dockerfile                   # Multi-stage Docker build
├── docker-compose.yml           # Docker Compose for local development
├── .env.example                 # Environment variables template
│
├── templates/                   # HTML templates
│   ├── home.html               # Home page
│   ├── about.html              # About page
│   ├── skills.html             # Skills showcase
│   ├── projects.html           # Projects portfolio
│   └── contact.html            # Contact page
│
├── static/                      # Static assets
│   ├── styles.css              # Main stylesheet
│   ├── resume.pdf              # Your resume
│   └── images/                 # Images and graphics
│
├── k8s/                         # Kubernetes manifests
│   ├── deployment.yaml         # Deployment, Service, HPA, RBAC
│   └── ingress.yaml            # Ingress, Network Policies, TLS
│
├── .github/
│   └── workflows/
│       └── ci-cd.yml           # GitHub Actions CI/CD pipeline
│
└── README.md                    # This file
```

## 🚀 Quick Start

### Prerequisites
- Go 1.22+
- Docker & Docker Compose
- kubectl (for Kubernetes deployment)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/YOUR_GITHUB/devops-portfolio.git
   cd devops-portfolio
   ```

2. **Run locally with Go**
   ```bash
   go run main.go
   ```
   
   Application will be available at `http://localhost:8080`

3. **Or run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

4. **Run tests**
   ```bash
   go test -v ./...
   ```

## 🐳 Docker Deployment

### Build the image
```bash
docker build -t devops-portfolio:latest .
```

### Run the container
```bash
docker run -p 8080:8080 devops-portfolio:latest
```

### Using Docker Compose
```bash
docker-compose up -d
docker-compose logs -f portfolio
```

### Push to Docker Hub
```bash
docker tag devops-portfolio:latest YOUR_GITHUB/devops-portfolio:latest
docker push YOUR_GITHUB/devops-portfolio:latest
```

## ☸️ Kubernetes Deployment

### Prerequisites
- kubectl configured and connected to cluster
- NGINX Ingress Controller installed
- cert-manager for TLS certificates (optional)

### Deploy to Kubernetes

1. **Update the deployment image**
   Edit `k8s/deployment.yaml` and replace `YOUR_GITHUB` with your Docker username:
   ```yaml
   image: YOUR_GITHUB/devops-portfolio:latest
   ```

2. **Deploy to cluster**
   ```bash
   # Deploy all manifests
   kubectl apply -f k8s/
   
   # Or deploy individually
   kubectl apply -f k8s/deployment.yaml
   kubectl apply -f k8s/ingress.yaml
   ```

3. **Verify deployment**
   ```bash
   # Check pods
   kubectl get pods -n portfolio
   
   # Check services
   kubectl get svc -n portfolio
   
   # Check ingress
   kubectl get ingress -n portfolio
   
   # View logs
   kubectl logs -n portfolio -l app=portfolio -f
   ```

4. **Access the application**
   ```bash
   # Port forward to local
   kubectl port-forward -n portfolio svc/portfolio 8080:80
   
   # Access at http://localhost:8080
   ```

### Update deployment
```bash
# Update image
kubectl set image deployment/portfolio \
  portfolio=YOUR_GITHUB/devops-portfolio:v1.1.0 \
  -n portfolio --record

# Verify rollout
kubectl rollout status deployment/portfolio -n portfolio
```

## 🔧 Configuration

### Environment Variables

Copy `.env.example` to `.env` and update with your details:

```bash
cp .env.example .env
```

Key variables:
- `PORT`: Application port (default: 8080)
- `YOUR_NAME`: Your name
- `YOUR_EMAIL`: Your email address
- `YOUR_GITHUB`: Your GitHub username
- `YOUR_LINKEDIN`: Your LinkedIn profile
- `YOUR_HANDLE`: Your social media handle

## 📝 Customization

### Update Personal Information
Replace all `YOUR_*` placeholders in:
- HTML templates (`templates/`)
- `.env.example`
- Kubernetes manifests (`k8s/deployment.yaml` and `k8s/ingress.yaml`)
- GitHub Actions workflow (`.github/workflows/ci-cd.yml`)

### Add Your Projects
Edit `templates/projects.html` with your actual projects, descriptions, and links.

### Customize Skills
Edit `templates/skills.html` to showcase your specific technologies and expertise.

### Update Resume
Replace `static/resume.pdf` with your actual resume.

### Customize Styling
Edit `static/styles.css` to match your preferred color scheme and design.

## 🔐 Security Features

- ✅ Non-root container user
- ✅ Read-only root filesystem
- ✅ Resource limits and requests
- ✅ Network policies
- ✅ RBAC with ServiceAccount
- ✅ Pod security policies
- ✅ TLS/HTTPS support
- ✅ Liveness and readiness probes
- ✅ Vulnerability scanning (Trivy)

## 📊 CI/CD Pipeline

The GitHub Actions workflow includes:

1. **Code Quality**: Go fmt, vet, tests
2. **Security Scanning**: Trivy filesystem scan
3. **Docker Build**: Multi-stage build, push to registry
4. **Container Scanning**: Trivy container image scan
5. **Kubernetes Deploy**: Rolling update to K8s cluster
6. **Health Checks**: Verify deployment health

Triggers:
- Push to `main` or `develop`
- Pull requests
- Version tags (`v*`)

## 📈 Monitoring & Observability

### Health Check
```bash
curl http://localhost:8080/health
```

### Kubernetes Health Checks
The deployment includes:
- **Liveness probe**: Restart unhealthy pods
- **Readiness probe**: Traffic routing only to ready pods
- **Startup probe**: Grace period for initialization

### Metrics
Enable Prometheus monitoring by uncommenting services in `docker-compose.yml`:
- Prometheus: http://localhost:9090
- Grafana: http://localhost:3000

## 🧪 Testing

### Run unit tests
```bash
go test -v ./...
```

### Run tests with coverage
```bash
go test -v -coverprofile=coverage.txt ./...
go tool cover -html=coverage.txt
```

### Test the API
```bash
# Health check
curl http://localhost:8080/health

# Test pages
curl http://localhost:8080/
curl http://localhost:8080/about
curl http://localhost:8080/skills
curl http://localhost:8080/projects
curl http://localhost:8080/contact
```

## 🌐 Deployment Options

### Local Machine
```bash
go run main.go
```

### Docker
```bash
docker run -p 8080:8080 devops-portfolio:latest
```

### Docker Compose
```bash
docker-compose up
```

### Kubernetes (Single Node)
```bash
kubectl apply -f k8s/
kubectl port-forward svc/portfolio 8080:80 -n portfolio
```

### Kubernetes (Multi-Node with Ingress)
```bash
kubectl apply -f k8s/
# Update DNS to point to ingress IP
# Access via portfolio.example.com
```

### Cloud Platforms
- **AWS**: Deploy to EKS using `k8s/` manifests
- **Azure**: Deploy to AKS using `k8s/` manifests
- **GCP**: Deploy to GKE using `k8s/` manifests

## 📚 Architecture

```
┌─────────────────────────────────────────┐
│        End User / Browser               │
└──────────────────┬──────────────────────┘
                   │
        ┌──────────▼──────────┐
        │   NGINX Ingress     │
        │  (TLS Termination)  │
        └──────────┬──────────┘
                   │
        ┌──────────▼──────────┐
        │  Kubernetes Service │
        │  (Load Balancer)    │
        └──────────┬──────────┘
                   │
    ┌──────────────┼──────────────┐
    │              │              │
┌───▼──┐      ┌───▼──┐      ┌───▼──┐
│ Pod1 │      │ Pod2 │ ...  │ PodN │
│  Go  │      │  Go  │      │  Go  │
└──────┘      └──────┘      └──────┘
   │              │              │
   └──────────────┼──────────────┘
                  │
        ┌─────────▼─────────┐
        │  Horizontal Pod   │
        │   Autoscaler      │
        │  (CPU/Memory)     │
        └───────────────────┘
```

## 🚨 Troubleshooting

### Container won't start
```bash
# Check logs
docker logs devops-portfolio

# Check permissions
docker run -it devops-portfolio:latest /bin/sh
```

### Kubernetes pod not running
```bash
# Describe pod for events
kubectl describe pod -n portfolio -l app=portfolio

# View logs
kubectl logs -n portfolio -l app=portfolio

# Check resource availability
kubectl top nodes
kubectl top pods -n portfolio
```

### Health check failing
```bash
# Test health endpoint
kubectl exec -it -n portfolio POD_NAME -- wget http://localhost:8080/health -O -
```

## 📄 License

This project is licensed under the MIT License. See LICENSE file for details.

## 👤 About

Built with ❤️ by YOUR_NAME

- **GitHub**: https://github.com/YOUR_GITHUB
- **LinkedIn**: https://linkedin.com/in/YOUR_LINKEDIN
- **Email**: YOUR_EMAIL@example.com

## 🤝 Contributing

This is a personal portfolio project, but suggestions and improvements are welcome!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Commit your changes (`git commit -am 'Add improvement'`)
4. Push to the branch (`git push origin feature/improvement`)
5. Open a Pull Request

## 📞 Contact

Have questions or opportunities? Let's connect!

- 📧 Email: YOUR_EMAIL@example.com
- 💼 LinkedIn: https://linkedin.com/in/YOUR_LINKEDIN
- 🐙 GitHub: https://github.com/YOUR_GITHUB
- 🌐 Portfolio: https://portfolio.example.com

---

**Built for DevOps Engineers, by DevOps Engineers** 🚀

<img width="1837" height="822" alt="image" src="https://github.com/user-attachments/assets/15396ca4-782d-44be-9736-404f0ed18079" />

