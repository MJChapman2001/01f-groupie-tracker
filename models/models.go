package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	//"strings"
)

func ReadJson(link, id string) []byte {
	response, err := http.Get(link+id)
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

type Locations struct {
	Id int `json:"id"`
	Locations []string `json:"locations"`
	Dates interface{} `json:"dates"`
}

func (l *Locations) New(id string) {
	data := ReadJson("https://groupietrackers.herokuapp.com/api/locations/", id)

	json.Unmarshal(data, &l)
}

type ConcertDates struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
}

func (cd *ConcertDates) New(id string) {
	data := ReadJson("https://groupietrackers.herokuapp.com/api/dates/", id)

	json.Unmarshal(data, &cd)
}

type Relations struct {
	Id int `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func (r *Relations) New(id string) {
	data := ReadJson("https://groupietrackers.herokuapp.com/api/relation/", id)

	json.Unmarshal(data, &r)

	// for i, j := range r.DatesLocations {
	// 	i = strings.Replace(i, "_", " ", -1)
	// 	i = strings.Replace(i, "-", ", ", -1)

	// 	for _, x := range j {
	// 		x = strings.Replace(x, "_", " ", -1)
	// 		x = strings.Replace(x, "-", ", ", -1)
	// 	}
	// }
}

type Artist struct {
	Id int `json:"id"`
	Image string `json:"image"`
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int `json:"creationDate"`
	FirstAlbum string `json:"firstAlbum"`
	Locations interface{} `json:"locations"`
	Dates interface{} `json:"concertDates"`
	Relations interface{} `json:"relations"`
}

func (a *Artist) New(id string) {
	data := ReadJson("https://groupietrackers.herokuapp.com/api/artists/", id)

	json.Unmarshal(data, &a)
}