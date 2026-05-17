// DevOps Portfolio Application Tests
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHealthEndpoint verifies the health check endpoint
func TestHealthEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("health check returned wrong status: got %v want %v",
			status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("health check returned unexpected content type: got %v want application/json",
			contentType)
	}

	expected := `{"status": "healthy"}`
	if rr.Body.String() != expected {
		t.Errorf("health check returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestHomePageHandler verifies the home page is served
func TestHomePageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("home.html")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("home page returned wrong status: got %v want %v",
			status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "text/html; charset=utf-8" {
		t.Errorf("home page returned unexpected content type: got %v want text/html",
			contentType)
	}
}

// TestAboutPageHandler verifies the about page is served
func TestAboutPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("about.html")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("about page returned wrong status: got %v want %v",
			status, http.StatusOK)
	}
}

// TestSkillsPageHandler verifies the skills page is served
func TestSkillsPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/skills", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("skills.html")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("skills page returned wrong status: got %v want %v",
			status, http.StatusOK)
	}
}

// TestProjectsPageHandler verifies the projects page is served
func TestProjectsPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/projects", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("projects.html")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("projects page returned wrong status: got %v want %v",
			status, http.StatusOK)
	}
}

// TestContactPageHandler verifies the contact page is served
func TestContactPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/contact", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := pageHandler("contact.html")
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("contact page returned wrong status: got %v want %v",
			status, http.StatusOK)
	}
}

// BenchmarkHealthCheckHandler benchmarks the health check endpoint
func BenchmarkHealthCheckHandler(b *testing.B) {
	handler := http.HandlerFunc(healthCheckHandler)
	rr := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/health", nil)
		handler.ServeHTTP(rr, req)
	}
}
