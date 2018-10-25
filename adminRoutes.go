package main

import (
	"encoding/json"
	"net/http"
)

func AdminTrackCount(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	count := countTrack()

	json.NewEncoder(w).Encode(count)
}

func AdminTracksDelete(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	//Return the struct as a json object.

	deleteTrackCollection()
	lastID = 0
	var emptyList []int
	trackID = emptyList

	json.NewEncoder(w).Encode("Not yet Implemented")
}
