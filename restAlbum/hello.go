package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Album struct {
	ID          string `json:"id"`
	ImageNumber string `json:"image_number"`
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	Year        string `json:"year"`
}

var albums []Album

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(albums)

}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range albums {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {

	albums = append(albums, Album{ID: "1", ImageNumber: "1", Name: "The Color and The Shape ", Artist: "Foo Fighters", Year: "1997"},
		Album{ID: "2", ImageNumber: "2", Name: "The Satanist ", Artist: "Behemoth", Year: "2014"})

	// initialize router
	router := mux.NewRouter()

	//endpoint
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbum).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
