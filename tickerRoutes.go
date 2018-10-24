package main

import (
	"encoding/json"
	"net/http"
)

var paging = 5 //Paging

func TickerLatest(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}

func Ticker(w http.ResponseWriter, r *http.Request) {

	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}

func TickerTimestamp(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode("Not yet Implemented")
}
