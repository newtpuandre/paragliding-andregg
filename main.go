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
	//IGC Routes located in igcRoutes.go
	r.HandleFunc("/paragliding/api", APIInfoRoute).Methods("GET")
	// /paragliding/ --> /paragliding/api
	r.HandleFunc("/paragliding/api/track", TrackIDPost).Methods("POST")
	r.HandleFunc("/paragliding/api/track", TrackIDAll).Methods("GET")
	r.HandleFunc("/paragliding/api/track/{igcId}", TrackID).Methods("GET")
	r.HandleFunc("/paragliding/api/track/{igcId}/{igcField}", TrackField).Methods("GET")

	//Ticker Routes located in tickerRoutes.go
	r.HandleFunc("/paragliding/api/ticker/latest", TickerLatest).Methods("GET")
	r.HandleFunc("/paragliding/api/ticker", Ticker).Methods("GET")
	r.HandleFunc("/paragliding/api/ticker/{timestamp}", TickerTimestamp).Methods("GET")

	//Webhook Routes located in webhookRoutes.go
	r.HandleFunc("/paragliding/api/webhook/new_track", WebhookNewTrack).Methods("POST")
	r.HandleFunc("/paragliding/api/webhook/new_track/{webhook_id}", WebhookIDGet).Methods("GET")
	r.HandleFunc("/paragliding/api/webhook/new_track/{webhook_id}", WebhookIDDelete).Methods("DELETE")

	//Admin Routes located in adminRoutes.go
	r.HandleFunc("/admin/api/tracks_count", AdminTrackCount).Methods("GET")
	r.HandleFunc("/admin/api/tracks", AdminTracksDelete).Methods("DELETE")
}

func main() {

	/*addr, err := determineListenAddress() //Get listening address
	if err != nil {
		log.Fatal(err)
	}*/

	//Setup router
	router := mux.NewRouter().StrictSlash(true)

	//Set up credentials for MongoDB
	dbInit()

	//Update ID counts from DB
	updateIdFromDB()

	//Make the router handle routes. Routes located in routes.go
	addRoutes(router)

	//dbInit()
	//Log fatal errors and start the server
	//log.Fatal(http.ListenAndServe(addr, router))
	log.Fatal(http.ListenAndServe(":8080", router))
}
