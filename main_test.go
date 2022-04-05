package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/api"
)

func Test_home(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HomeHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()
	response, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, response.StatusCode)
	}
}

func Test_artists(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/artists", api.ArtistHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()
	t.Run("not found", func(t *testing.T) {
		response, err := http.Get(ts.URL + "/artists?id=0")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if response.StatusCode != http.StatusNotFound {
			t.Errorf("Expected %d, received %d", http.StatusNotFound, response.StatusCode)
		}
	})
	t.Run("found", func(t *testing.T) {
		response, err := http.Get(ts.URL + "/artists?id=1")
		if err != nil {
			t.Errorf("Expected nil, received %s", err.Error())
		}
		if response.StatusCode != http.StatusOK {
			t.Errorf("Expected %d, received %d", http.StatusOK, response.StatusCode)
		}
	})
}

func Test_locations(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/locations", api.LocationHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()
	response, err := http.Get(ts.URL + "/locations")
	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, response.StatusCode)
	}
}