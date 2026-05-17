// DevOps Portfolio Application
// A modern personal developer portfolio + DevOps showcase platform
// Built with Go, Docker, Kubernetes, and Cloud-native technologies
package main

import (
	"log"
	"net/http"
	"os"
)

// PageHandler serves HTML pages from the templates directory
func pageHandler(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "templates/"+filename)
	}
}

// healthCheckHandler provides a health check endpoint for Kubernetes liveness probes
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}

// getPort retrieves the port from environment variable or defaults to 8080
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}

func main() {
	// Serve static assets (CSS, images, resume, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route handlers for pages
	http.HandleFunc("/", pageHandler("home.html"))
	http.HandleFunc("/home", pageHandler("home.html"))
	http.HandleFunc("/about", pageHandler("about.html"))
	http.HandleFunc("/skills", pageHandler("skills.html"))
	http.HandleFunc("/projects", pageHandler("projects.html"))
	http.HandleFunc("/contact", pageHandler("contact.html"))

	// Health check endpoint for container orchestration
	http.HandleFunc("/health", healthCheckHandler)

	// Log server startup
	port := getPort()
	log.Printf("🚀 DevOps Portfolio Server starting on http://localhost%s", port)
	log.Printf("📖 Navigate to http://localhost%s/ to view the portfolio", port)

	// Start HTTP server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("❌ Server error: %v", err)
	}
}
