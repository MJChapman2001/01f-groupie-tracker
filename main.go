package main

import (
	"log"
	"net/http"

	"groupie-tracker/api"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HomeHandler)
	mux.HandleFunc("/artists", api.ArtistHandler)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
