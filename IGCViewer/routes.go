package IGCViewer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/marni/goigc" //Parse and handle IGC Files

	"github.com/rickb777/date/period" //Used to convert to ISO-8601 Duration Format

	"github.com/gorilla/mux" //Router
)

//store startime of the application for further use
var startTime time.Time = time.Now()

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

//IgcIdPost handles and adds URL of flight routes into memory
func IgcIdPost(w http.ResponseWriter, r *http.Request) {
	//TODO: Handle ERRORS

	//Decode incoming url
	var decodedUrl Url
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&decodedUrl)

	if err != nil {
		//Handle error
		panic(err)
	}

	//Parse IGC File from URL
	s := decodedUrl.Url

	//Handle if the URL isnt a .igc file
	track, err := igc.ParseLocation(s)
	if err != nil {
		fmt.Errorf("Problem reading the track", err)
	}

	//Fill Track struct with required information
	var newTrack Track

	newTrack.Pilot = track.Pilot
	newTrack.Glider = track.GliderType
	newTrack.Glider_id = track.GliderID
	newTrack.Track_length = len(track.Points) //This is probably wrong. Check!
	newTrack.H_date = track.Date.String()

	tracks = append(tracks, newTrack)

	//Add ID to array over used ids
	trackID = append(trackID, lastID)
	var idStruct Url_ID
	idStruct.Id = lastID
	lastID++

	//Return the struct as a json object.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(idStruct)

}

//IgcIdAll returns an json array with all track ids
func IgcIdAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if len(trackID) <= 0 { //Do we have content to display?

		emptyArray := make([]string, 0) //Create an empty array and return it.
		json.NewEncoder(w).Encode(emptyArray)

	} else { //Show array in memory

		json.NewEncoder(w).Encode(trackID)

	}

}

//IgcId returns a json object with a specific id
func IgcId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	igcId := vars["igcId"]

	fmt.Println(igcId)
	//Check if the parameter passed is an integer.
	i, err := strconv.Atoi(igcId)

	if err == nil && i <= len(tracks) { //Is an int and not bigger than tracks in memory
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(tracks[i])
	} else {
		//Return bad request
		panic(err)
	}

}

//IgcField returns information about a specific field from a track
func IgcField(w http.ResponseWriter, r *http.Request) {

}
