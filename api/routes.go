package api

import (
	"html/template"
	"net/http"
	"strings"
	"strconv"
	//"fmt"
	
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

	relations := make(map[string][]string)

	for i, j := range Artistrelations.DatesLocations {
		var dates []string
		i = strings.Title(strings.Replace(i, "-", ", ", -1))
		i = strings.Title(strings.Replace(i, "_", " ", -1))
		for _, x := range j {
			x = strings.Replace(x, "-", "/", -1)
			dates = append(dates, x)
		}
		relations[i] = dates
	}

	Artist.Relations = relations

	if err := tmpl.Execute(w, Artist); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/locations.html"))

	Locations := make(map[string]map[string][]models.Artist)

	LocationsData := LoadAllLocations()

	for _, i := range LocationsData {
		var Artist models.Artist

		id := strconv.Itoa(i.Id)
		Artist.New(id)

		for _, j := range i.Locations {
			split := strings.Split(j, "-")
			group := strings.Replace(split[1], "_", " ", -1)

			if group == "uk" || group == "usa" {
				group = strings.ToUpper(group)
			} else {
				group = strings.Title(group)
			}

			city := strings.Title(strings.Replace(split[0], "_", " ", -1))

			if Locations[group] == nil {
				Locations[group] = make(map[string][]models.Artist)
			}

			Locations[group][city] = append(Locations[group][city], Artist)
		}
	}

	if err := tmpl.Execute(w, Locations); err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}