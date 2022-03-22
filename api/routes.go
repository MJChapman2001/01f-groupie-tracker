package api

import (
	"html/template"
	"net/http"
	
	"groupie-tracker/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	if r.URL.Path != "/" {
		http.Error(w, "404 Status Not Found", http.StatusNotFound)
		return
	}

	Artists := LoadAllArtists()

	if err := tmpl.Execute(w, Artists); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/artist.html"))

	id := r.URL.Query().Get("id")

	var Artist models.Artist
	Artist.New(id)

	var Artistrelations models.Relations
	Artistrelations.New(id)
	Artist.Relations = Artistrelations

	if err := tmpl.Execute(w, Artist); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}