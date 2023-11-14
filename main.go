package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {

	//Compile Tailwind
	err := exec.Command("npx", "postcss", "static/tailwind.css", "-o", "static/output.css").Run()
	if err != nil {
		fmt.Println("Failed to compile CSS", http.StatusInternalServerError, err)
		return
	}

	// Serve static assets like CSS, JS, images, etc.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/index.html")
	})

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Content loaded with HTMX!"))
	})

	log.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
