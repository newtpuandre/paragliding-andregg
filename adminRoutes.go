package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//AdminTrackCount counts amount of track in database and return to get request
func AdminTrackCount(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	count := countTrack(&Credentials)

	err := json.NewEncoder(w).Encode(count)
	if err != nil {
		fmt.Println(err)
	}
}

//AdminTracksDelete deletes and resets track collection and variables in hook collection
func AdminTracksDelete(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	//Return the struct as a json object.

	count := deleteTrackCollection(&Credentials)

	lastID = 0
	var emptyList []int
	trackID = emptyList

	clearWebhookID(&Credentials)

	err := json.NewEncoder(w).Encode(count)
	if err != nil {
		fmt.Println(err)
	}
}
