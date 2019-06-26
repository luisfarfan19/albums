package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allAlbums, err := RetrieveAlbums()
	if err != nil {
		return
	}
	fmt.Println(allAlbums)
	json.NewEncoder(w).Encode(allAlbums)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	allAlbums, err := RetrieveAlbums()
	if err != nil {
		return
	}
	for _, item := range allAlbums {
		if string(item.Id) == params["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getAlbumByArtist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	allAlbums, err := RetrieveAlbumsByArtist(params["artist"])
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(allAlbums)
}

func main() {

	//a := Album{
	//	Id:       2,
	//	ArtistId: 1,
	//	Name:     "FOO FIGHTERS",
	//	Year:     1995,
	//	Price:    100,
	//}
	//
	//err := Create(a)
	//if err != nil{
	//	log.Fatal(err)
	//}

	// initialize router
	router := mux.NewRouter()

	//endpoint
	router.HandleFunc("/albums", getAlbums).Methods("GET")
	router.HandleFunc("/album/{id}", getAlbum).Methods("GET")
	router.HandleFunc("/album/artist/{artist}", getAlbumByArtist).Methods("GET")

	log.Fatal(http.ListenAndServe(":5000", router))
}
