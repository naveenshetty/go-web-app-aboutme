package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/home.html")
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Custom routes
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/whoIam", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/whoIam.html")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/about.html")
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/contact.html")
	})

	//log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
