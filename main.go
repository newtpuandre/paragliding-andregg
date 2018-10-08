package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux" //Router
)

func main() {

	//Setup router
	router := mux.NewRouter().StrictSlash(true)

	//Make the router handle routes. Routes located in routes.go
	router.HandleFunc("/igcinfo/api", APIInfoRoute).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc", IgcIdPost).Methods("POST")
	router.HandleFunc("/igcinfo/api/igc", IgcIdAll).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}", IgcId).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}/{igcField}", IgcField).Methods("GET")

	http.Handle("/", router)

	//Log fatal errors and start the server
	log.Fatal(http.ListenAndServe(":80", router))
}
