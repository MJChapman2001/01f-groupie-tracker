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
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}