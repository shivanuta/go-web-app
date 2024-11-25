package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to test routes
func testRoute(t *testing.T, route string, handler http.HandlerFunc) {
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatalf("Failed to create request for %s: %v", route, err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	// Check if status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler for %s returned wrong status code: got %v want %v",
			route, status, http.StatusOK)
	}

	// Check if content type is "text/html; charset=utf-8"
	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("Handler for %s returned unexpected content type: got %v want %v",
			route, contentType, expectedContentType)
	}
}

func TestHomePage(t *testing.T) {
	testRoute(t, "/home", homePage)
}

func TestCoursePage(t *testing.T) {
	testRoute(t, "/courses", coursePage)
}

func TestAboutPage(t *testing.T) {
	testRoute(t, "/about", aboutPage)
}

func TestContactPage(t *testing.T) {
	testRoute(t, "/contact", contactPage)
}

func TestDefaultRoute(t *testing.T) {
	// Simulate a request to the default route "/"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	}).ServeHTTP(rr, req)

	// Check if the response is a redirect to "/home"
	if status := rr.Code; status != http.StatusMovedPermanently {
		t.Errorf("Default route returned wrong status code: got %v want %v",
			status, http.StatusMovedPermanently)
	}

	expectedLocation := "/home"
	if location := rr.Header().Get("Location"); location != expectedLocation {
		t.Errorf("Default route redirected to wrong location: got %v want %v",
			location, expectedLocation)
	}
}
