package main

import (
	"log"
	"net/http"

	"./IGCViewer"
	"github.com/gorilla/mux" //Router
)

func main() {

	//Setup router
	router := mux.NewRouter().StrictSlash(true)

	//Make the router handle routes. Routes located in
	//IGCViewer/routes.go
	router.HandleFunc("/igcinfo/api", IGCViewer.APIInfoRoute).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc", IGCViewer.IgcIdPost).Methods("POST")
	router.HandleFunc("/igcinfo/api/igc", IGCViewer.IgcIdAll).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}", IGCViewer.IgcId).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}/{igcField}", IGCViewer.IgcField).Methods("GET")

	//Log fatal errors and start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
