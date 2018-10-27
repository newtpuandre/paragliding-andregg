package main

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/marni/goigc" //Parse and handle IGC Files

	"github.com/rickb777/date/period" //Used to convert to ISO-8601 Duration Format

	"github.com/gorilla/mux" //Router
)

//store startime of the application for further use
var startTime = time.Now()

//Track IDs stored in memory
var trackID []int
var lastID int

//APIInfoRoute returns a struct with information about the api
func APIInfoRoute(w http.ResponseWriter, r *http.Request) {

	//Fill the info struct with uptime and other information.
	trackerInfo := apiInfo{Uptime: period.Between(startTime, time.Now()), Info: "Service for Paragliding tracks.", Version: "v1"}

	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode(trackerInfo)
}

//TrackIDPost handles and adds URL and flight routes into memory
func TrackIDPost(w http.ResponseWriter, r *http.Request) {

	//Decode incoming url
	var decodedURL URL
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedURL)

	if err != nil {
		//Malformed content body.
		http.Error(w, "Malformed content body", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	//Parse IGC File from URL
	s := decodedURL.URL

	if !strings.Contains(s, ".igc") { //Not a secure way to check filetype...
		http.Error(w, "Not a IGC file", http.StatusBadRequest)
		return
	}

	track, err := igc.ParseLocation(s)
	if err != nil {
		//Bad IGC file or bad URL
		http.Error(w, "Bad file or URL", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	//Fill Track struct with required information
	var newTrack Track

	newTrack.Timestamp = time.Now().Unix()
	newTrack.Pilot = track.Pilot
	newTrack.Glider = track.GliderType
	newTrack.Glider_id = track.GliderID
	newTrack.H_date = track.Date.String()
	newTrack.Track_src_url = s

	//Distance calculation for a track
	for i := 0; i < len(track.Points)-1; i++ {
		newTrack.Track_length += track.Points[i].Distance(track.Points[i+1])
	}

	//Pass DB Credentials and new track to function
	insertTrack(&newTrack, &Credentials)

	//Add ID to array for used ids
	trackID = append(trackID, lastID)

	//Remember to count up used ids
	lastID++

	//Fill return struct
	var idStruct URLID
	idStruct.ID = lastID

	//Return the struct as a json object.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(idStruct)

	//Invoke webhooks
	invokeWebHook()

}

//TrackIDAll returns an json array with all track ids
func TrackIDAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if len(trackID) <= 0 { //Do we have content to display?

		emptyArray := make([]string, 0) //Create an empty array and return it.
		json.NewEncoder(w).Encode(emptyArray)

	} else { //Show array in memory

		json.NewEncoder(w).Encode(trackID)

	}

}

//TrackID returns a json object with a specific id
func TrackID(w http.ResponseWriter, r *http.Request) {
	//Get parameters
	var tracks = getAllTracks(&Credentials)
	vars := mux.Vars(r)
	igcID := vars["igcId"]

	//Check if the parameter passed is an integer.
	i, err := strconv.Atoi(igcID)

	//Absolute value the integer. We dont accept negative numbers!
	if i < 0 {
		i = i * -1
	}

	if err == nil && i <= len(tracks)-1 { //Is an int and not bigger than tracks in memory

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var tempTrack TrackNoTimestamp
		tempTrack.Glider = tracks[i].Glider
		tempTrack.Glider_id = tracks[i].Glider_id
		tempTrack.H_date = tracks[i].H_date
		tempTrack.Pilot = tracks[i].Pilot
		tempTrack.Track_length = tracks[i].Track_length
		tempTrack.Track_src_url = tracks[i].Track_src_url
		json.NewEncoder(w).Encode(tempTrack)

	} else {
		//Return bad request
		http.Error(w, "", 404) //404 Not found
	}

}

//TrackField returns information about a specific field from a track
func TrackField(w http.ResponseWriter, r *http.Request) {
	//Header is not set because it defaults to text/plain charset=utf-8
	var tracks = getAllTracks(&Credentials)
	//Get parameters
	vars := mux.Vars(r)
	igcID := vars["igcId"]
	igcField := vars["igcField"]

	//Make the first letter uppcase
	//to be the same as struct variables
	upperIgcFIeld := strings.Title(igcField)

	//Try to convert the paramter to an int
	i, err := strconv.Atoi(igcID)

	//Absolute value the integer. We dont accept negative numbers!
	if i < 0 {
		i = i * -1
	}

	if upperIgcFIeld == "Timestamp" {
		http.Error(w, "", 400)
		return
	}

	if err != nil || i >= len(tracks) { //Could not convert to int and

		http.Error(w, "", 404) //404 Not found
		return
	}

	track := tracks[i]

	//Try to match user input with field from the selected track struct
	ref := reflect.ValueOf(track)
	f := reflect.Indirect(ref).FieldByName(upperIgcFIeld)

	if strings.Contains(f.String(), "invalid Value") { //Does the field exist?
		//Return 404 when it doesn't exist

		http.Error(w, "", 404) //404 Not found
		return
	}

	//Handle type
	if strings.Contains(f.String(), "float64 Value") {
		json.NewEncoder(w).Encode(f.Float()) //Print as float
	} else {
		json.NewEncoder(w).Encode(f.String()) //Print as string
	}

}
