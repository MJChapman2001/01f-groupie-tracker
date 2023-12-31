package api

import (
	"log"
	"net/http"
	"encoding/json"
	"strings"
	"io"
	"io/ioutil"

	"groupie-tracker/models"
)

func LoadAllArtists() []models.Artist {
	var Artists []models.Artist

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := string(responseData)

	reader := strings.NewReader(data)
	dec := json.NewDecoder(reader)

	if err := dec.Decode(&Artists); err == io.EOF {
		
	} else if err != nil {
		log.Fatal(err)
	}

	return Artists
}