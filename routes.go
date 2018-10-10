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

//Tracks stored in memory
var trackID []int
var lastID int

//Uses the trackID index as a lookup indirectly.
//The two elements are not directly attached.
var tracks []Track

//APIInfoRoute returns a struct with information about the api
func APIInfoRoute(w http.ResponseWriter, r *http.Request) {

	//Fill the info struct with uptime and other information.
	trackerInfo := apiInfo{Uptime: period.Between(startTime, time.Now()), Info: "Service for IGC tracks.", Version: "v1"}

	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode(trackerInfo)
}

//IgcIDPost handles and adds URL and flight routes into memory
func IgcIDPost(w http.ResponseWriter, r *http.Request) {
	//TODO: Handle ERRORS

	//Decode incoming url
	var decodedURL Url
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedURL)

	if err != nil {
		//Malformed content body.
		http.Error(w, "Malformed content body", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	//Parse IGC File from URL
	s := decodedURL.Url

	//Handle if the URL isnt a .igc file
	track, err := igc.ParseLocation(s)
	if err != nil {
		//Bad IGC file or bad URL
		http.Error(w, "Bad file or URL", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	//Fill Track struct with required information
	var newTrack Track

	newTrack.Pilot = track.Pilot
	newTrack.Glider = track.GliderType
	newTrack.Glider_id = track.GliderID
	newTrack.H_date = track.Date.String()

	//Loop through points and add the distance between them together
	newTrack.Track_length = track.Points[0].Distance(track.Points[len(track.Points)-1])
	//for i := 0; i < len(track.Points)-1; i++ {
	//	newTrack.Track_length += track.Points[i].Distance(track.Points[i+1])
	//}

	//Add track to array for all tracks
	tracks = append(tracks, newTrack)

	//Add ID to array for used ids
	trackID = append(trackID, lastID)

	//Fill return struct
	var idStruct Url_ID
	idStruct.Id = lastID

	//Remember to count up used ids
	lastID++

	//Return the struct as a json object.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(idStruct)

}

//IgcIDAll returns an json array with all track ids
func IgcIDAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if len(trackID) <= 0 { //Do we have content to display?

		emptyArray := make([]string, 0) //Create an empty array and return it.
		json.NewEncoder(w).Encode(emptyArray)

	} else { //Show array in memory

		json.NewEncoder(w).Encode(trackID)

	}

}

//IgcID returns a json object with a specific id
func IgcID(w http.ResponseWriter, r *http.Request) {
	//Get parameters
	vars := mux.Vars(r)
	igcID := vars["igcId"]

	//Check if the parameter passed is an integer.
	i, err := strconv.Atoi(igcID)

	if err == nil && i < len(tracks) { //Is an int and not bigger than tracks in memory
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(tracks[i])
	} else {
		//Return bad request
		http.Error(w, "", 404) //404 Not found
	}

}

//IgcField returns information about a specific field from a track
func IgcField(w http.ResponseWriter, r *http.Request) {
	//Get parameters
	vars := mux.Vars(r)
	igcID := vars["igcId"]
	igcField := vars["igcField"]

	//Make the first letter uppcase
	//to be the same as struct variables
	upperIgcFIeld := strings.Title(igcField)

	//Try to convert the paramter to an int
	i, err := strconv.Atoi(igcID)

	if err == nil && i < len(tracks) { //Is an int and not bigger than tracks in memory
		track := tracks[i]

		//Try to match user input with field from the selected track struct
		r := reflect.ValueOf(track)
		f := reflect.Indirect(r).FieldByName(upperIgcFIeld)

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

	} else {
		http.Error(w, "", 404) //404 Not found
	}

}
