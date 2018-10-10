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

func addRoutes(r *mux.Router) {
	//Make the router handle routes. Routes located in routes.go
	r.HandleFunc("/igcinfo/api", APIInfoRoute).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc", IgcIDPost).Methods("POST")
	r.HandleFunc("/igcinfo/api/igc", IgcIDAll).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc/{igcId}", IgcID).Methods("GET")
	r.HandleFunc("/igcinfo/api/igc/{igcId}/{igcField}", IgcField).Methods("GET")
}

func main() {

	addr, err := determineListenAddress() //Get listening address
	if err != nil {
		log.Fatal(err)
	}

	//Setup router
	router := mux.NewRouter().StrictSlash(true)

	//Make the router handle routes. Routes located in routes.go
	addRoutes(router)

	//Log fatal errors and start the server
	log.Fatal(http.ListenAndServe(addr, router))
}
