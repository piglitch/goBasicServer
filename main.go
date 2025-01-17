package main

import (	
	"net/http"
	"log"
)

func main() {
	// fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/texts/new", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/newtext.html")
	})

	log.Println("Server is running!")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}