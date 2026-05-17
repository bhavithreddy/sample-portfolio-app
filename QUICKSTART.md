# Quick Start Guide

Get your DevOps portfolio up and running in minutes!

## 5-Minute Setup

### 1. Personalize Your Portfolio

Edit these placeholders with your information:

```bash
# Find and replace in all files
find . -type f \( -name "*.html" -o -name "*.md" -o -name "*.yml" -o -name "*.yaml" \) \
  -exec sed -i 's/YOUR_NAME/Your Name/g' {} \;
  
sed -i 's/YOUR_EMAIL/your.email@example.com/g' $(find . -type f)
sed -i 's/YOUR_GITHUB/your-github-username/g' $(find . -type f)
sed -i 's/YOUR_LINKEDIN/your-linkedin-profile/g' $(find . -type f)
sed -i 's/YOUR_HANDLE/your-twitter-handle/g' $(find . -type f)
```

Or manually edit:
- `templates/*.html` - Replace placeholders
- `.env.example` - Your contact info
- `k8s/deployment.yaml` - Docker image
- `README.md` - Links and info

### 2. Add Your Projects

Edit `templates/projects.html` and replace the example projects with yours.

### 3. Update Skills

Edit `templates/skills.html` to showcase your expertise.

### 4. Add Your Resume

Replace `static/resume.pdf` with your actual resume.

## Run Locally

### Option A: Direct Go

```bash
go run main.go
# Access: http://localhost:8080
```

### Option B: Docker

```bash
docker build -t portfolio:latest .
docker run -p 8080:8080 portfolio:latest
# Access: http://localhost:8080
```

### Option C: Docker Compose

```bash
docker-compose up
# Access: http://localhost:8080
```

## Deploy to Production

### Option 1: Kubernetes

```bash
# Push Docker image
docker build -t YOUR_GITHUB/devops-portfolio:latest .
docker push YOUR_GITHUB/devops-portfolio:latest

# Update k8s/deployment.yaml with your image
sed -i 's|image:.*|image: YOUR_GITHUB/devops-portfolio:latest|' k8s/deployment.yaml

# Deploy
kubectl apply -f k8s/

# Access
kubectl port-forward svc/portfolio 8080:80 -n portfolio
```

### Option 2: Docker Compose (Single Server)

```bash
# Copy files to server
scp -r . user@server:/opt/portfolio/

# On server
cd /opt/portfolio
docker-compose up -d

# Access: http://server-ip:8080
```

### Option 3: Traditional (Linux Server)

```bash
# Copy files
scp -r . user@server:/opt/portfolio/

# On server
cd /opt/portfolio
go build -o portfolio
./portfolio &
```

## Common Commands

### Development

```bash
# Run tests
go test -v ./...

# Format code
go fmt ./...

# Lint
go vet ./...

# Build
go build -o portfolio

# Using Make
make test
make fmt
make build
```

### Docker

```bash
# Build
docker build -t portfolio:latest .

# Run
docker run -p 8080:8080 portfolio:latest

# Using Make
make docker-build
make docker-run
```

### Kubernetes

```bash
# Deploy
kubectl apply -f k8s/

# Check status
kubectl get pods -n portfolio
kubectl logs -n portfolio -l app=portfolio

# Port forward
kubectl port-forward -n portfolio svc/portfolio 8080:80

# Using Make
make k8s-deploy
make k8s-status
make k8s-logs
```

## Troubleshooting

### Port already in use

```bash
# Change port
PORT=9000 go run main.go

# Or kill existing process
lsof -ti:8080 | xargs kill -9
```

### File not found errors

```bash
# Ensure directory structure
templates/
static/
  ├── styles.css
  └── resume.pdf
```

### Docker build fails

```bash
# Clear cache
docker build --no-cache -t portfolio:latest .

# Check Go version
go version  # Should be 1.22+
```

### Kubernetes pod fails

```bash
# Check pod status
kubectl describe pod -n portfolio <POD_NAME>

# View logs
kubectl logs -n portfolio <POD_NAME>

# Check events
kubectl get events -n portfolio
```

## Next Steps

1. ✅ Customize with your information
2. ✅ Add your projects and skills
3. ✅ Test locally
4. ✅ Build Docker image
5. ✅ Set up domain (optional)
6. ✅ Deploy to production
7. ✅ Configure HTTPS/TLS (optional)
8. ✅ Set up monitoring (optional)
9. ✅ Share with others!

## Useful Resources

- [README.md](README.md) - Full documentation
- [DEPLOYMENT.md](DEPLOYMENT.md) - Detailed deployment guide
- [ARCHITECTURE.md](ARCHITECTURE.md) - Technical architecture
- [Go Documentation](https://go.dev/doc/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [Docker Documentation](https://docs.docker.com/)

## Getting Help

### Common Issues

**Application not responding**
```bash
curl -v http://localhost:8080/health
```

**Check logs**
```bash
# Docker
docker logs -f devops-portfolio

# Kubernetes
kubectl logs -f -n portfolio deployment/portfolio

# Local
# Check console output
```

**File permissions**
```bash
chmod 755 portfolio
chmod 644 static/*
chmod 644 templates/*
```

## Customization Examples

### Change color scheme

Edit `static/styles.css`:
```css
--accent-blue: #0ea5e9;
--accent-purple: #a855f7;
--accent-cyan: #06b6d4;
```

### Add new page

1. Create `templates/newpage.html`
2. Add route in `main.go`:
   ```go
   http.HandleFunc("/newpage", pageHandler("newpage.html"))
   ```
3. Add link in `templates/home.html`:
   ```html
   <li><a href="/newpage">New Page</a></li>
   ```

### Change application port

```bash
# Local
PORT=3000 go run main.go

# Docker
docker run -p 3000:8080 -e PORT=8080 portfolio:latest

# Kubernetes
kubectl set env deployment/portfolio PORT=3000 -n portfolio
```

## Performance Tips

- Keep images under 5MB each
- Minimize CSS/JavaScript
- Use CDN for static assets (optional)
- Enable gzip compression
- Set appropriate caching headers
- Regular backups (Git)

## Security Checklist

- [ ] All `YOUR_*` placeholders replaced
- [ ] No secrets in code or images
- [ ] HTTPS enabled (use Let's Encrypt)
- [ ] Security headers configured
- [ ] Container image scanned (Trivy)
- [ ] Access logs monitored
- [ ] Email alerts configured
- [ ] Backups tested

---

**Congratulations!** Your DevOps portfolio is ready. Show the world your skills! 🚀