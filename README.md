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



## 🧪 Testing

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


## 📄 License

This project is licensed under the MIT License. See LICENSE file for details.



---

**Built for DevOps Engineers, by DevOps Engineers** 🚀

<img width="1837" height="822" alt="image" src="https://github.com/user-attachments/assets/15396ca4-782d-44be-9736-404f0ed18079" />

