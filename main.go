package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux" //Router
)

func determineListenAddress() (string, error) { //Inorder to get the port heroku assigns us
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {

	addr, err := determineListenAddress() //Get listening address
	if err != nil {
		log.Fatal(err)
	}

	//Setup router
	router := mux.NewRouter().StrictSlash(true)

	//Make the router handle routes. Routes located in routes.go
	router.HandleFunc("/igcinfo/api", APIInfoRoute).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc", IgcIDPost).Methods("POST")
	router.HandleFunc("/igcinfo/api/igc", IgcIDAll).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}", IgcID).Methods("GET")
	router.HandleFunc("/igcinfo/api/igc/{igcId}/{igcField}", IgcField).Methods("GET")

	//Log fatal errors and start the server
	log.Fatal(http.ListenAndServe(addr, router))
}
