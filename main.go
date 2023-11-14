package main

import (
	"log"
	"net/http"
)

func main() {

	// Serve static assets like CSS, JS, images, etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Content loaded with HTMX!"))
	})

	log.Println("Server is running")
	http.ListenAndServe(":8080", nil)
}
