package main

import (
	"log"
	"net/http"
	"os"
)

// Serve a file with error handling for non-existent files
func serveFile(w http.ResponseWriter, r *http.Request, filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "404 - File Not Found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}

// Handlers for different routes
func homePage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/contact.html")
}

func main() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Default route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusMovedPermanently)
	})

	// Serve static assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
