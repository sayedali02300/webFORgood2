package main

import (
	groupie "GROUPIE/funcs"
	"log"
	"net/http"
)

func main() {
	groupie.Init()
	mux := http.NewServeMux()

	mux.Handle("/pics/", http.StripPrefix("/pics/", http.FileServer(http.Dir("pics"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	port := "8989"
	mux.HandleFunc("/", groupie.HomePage)
	mux.HandleFunc("/ArtistDetail", groupie.PostPage)
	mux.HandleFunc("/suggest", groupie.SuggestHandler)
	notFound := groupie.With404Handler(mux)

	log.Printf("Server running on http://localhost:%s", port)

	err := http.ListenAndServe(":"+port, notFound)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
