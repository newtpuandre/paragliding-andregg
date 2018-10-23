package main

import (
	"encoding/json"
	"net/http"
)

func WebhookNewTrack(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}

func WebhookIDGet(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}

func WebhookIDDelete(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}
