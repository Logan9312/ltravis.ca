package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Use `PORT` provided in environment or default to 3000
	port := envPortOr("3000")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Content loaded with HTMX!"))
	})

	log.Println("Server is running")
	log.Fatal(http.ListenAndServe(port, nil))
}

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
