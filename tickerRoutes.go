package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var paging = 5 //Paging

//TickerLatest returns the latest timestamp to the get request
func TickerLatest(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var tempTracks = getAllTracks(&Credentials)
	count := len(tempTracks) - 1

	if count < 0 {
		http.Error(w, "Nothing to show", 400) //400 bad request
		return
	}

	//Return the struct as a json object.
	err := json.NewEncoder(w).Encode(tempTracks[count].Timestamp)
	if err != nil {
		fmt.Println(err)
	}
}

//Ticker returns information to the get request about the last 5 tracks inserted
func Ticker(w http.ResponseWriter, r *http.Request) {
	start := time.Now() //function processing time

	var tempTracks = getAllTracks(&Credentials)

	//If there are no tracks return error
	var trackCount = len(tempTracks) - 1
	if trackCount < 0 {
		http.Error(w, "No tracks available", 400) //400 bad request
		return
	}
	var tempTicker tickerStruct

	tempTicker.TLatest = tempTracks[trackCount].Timestamp
	tempTicker.TStart = tempTracks[0].Timestamp

	//Stop value is the id of the last track in response
	var stop = paging - 1

	//If there are less than five tracks, use the latest as stop
	if trackCount < 5 {
		stop = trackCount
	}

	//Add Track ids to array
	for i := 0; i <= stop; i++ {
		tempTicker.Tracks = append(tempTicker.Tracks, i)
	}

	//Get timestamp of the last id.
	tempTicker.TStop = tempTracks[tempTicker.Tracks[len(tempTicker.Tracks)-1]].Timestamp

	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Function run time
	tempTicker.Processing = time.Since(start) / 1000000 //Convert to ms
	//Return the struct as a json object.
	err := json.NewEncoder(w).Encode(tempTicker)
	if err != nil {
		fmt.Println(err)
	}
}

//TickerTimestamp returns tracks bigger than provided timestamp to the post request
func TickerTimestamp(w http.ResponseWriter, r *http.Request) {
	//Get parameters
	vars := mux.Vars(r)
	timeStamp := vars["timestamp"]

	//Check if the parameter passed is an integer.
	newTimeStamp, err := strconv.Atoi(timeStamp)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "", 400) //400 bad request
		return
	}

	start := time.Now() //Function processing time
	var tempTracks = getAllTracks(&Credentials)

	//If there are no tracks return error
	var trackCount = len(tempTracks) - 1
	if trackCount < 0 {
		http.Error(w, "No tracks available", 400) //400 bad request
		return
	}

	//Find first id with bigger timestamp
	var index = -1
	for i := range tempTracks {
		if tempTracks[i].Timestamp > int64(newTimeStamp) {
			index = i
			break
		}
	}

	//If there are no bigger timestamp than provided. Error
	if index == -1 {
		http.Error(w, "No timestamp is bigger", 400)
		return
	}

	var tempTicker tickerStruct

	tempTicker.TLatest = tempTracks[trackCount].Timestamp
	tempTicker.TStart = tempTracks[index].Timestamp

	//Stop value is the id of the last track in response
	var stop = index + paging

	//Are there less than 5 pages?
	if index+paging > trackCount {
		stop = trackCount
	}

	//Add ids to an array
	for i := index; i <= stop; i++ {
		tempTicker.Tracks = append(tempTicker.Tracks, i)
	}

	//Get timestamp of the last id.
	tempTicker.TStop = tempTracks[tempTicker.Tracks[len(tempTicker.Tracks)-1]].Timestamp //Might be the correct behavior?
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Function run time
	tempTicker.Processing = time.Since(start) //Convert to ms
	//Return the struct as a json object.
	err = json.NewEncoder(w).Encode(tempTicker)
	if err != nil {
		fmt.Println(err)
	}
}
