package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Flytte til Database
var tickerList []tickerStruct

//stored in memory
var tickerID []int
var lastTickerID int

var paging = 5 //Paging

func TickerLatest(w http.ResponseWriter, r *http.Request) {
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var tempTracks = getAllTracks()
	count := len(tempTracks) - 1

	//Return the struct as a json object.
	json.NewEncoder(w).Encode(tempTracks[count].Timestamp)
}

func Ticker(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var tempTracks = getAllTracks()

	var trackCount = len(tempTracks) - 1
	var tempTicker tickerStruct

	tempTicker.TLatest = tempTracks[trackCount].Timestamp
	tempTicker.TStart = tempTracks[0].Timestamp

	var stop = paging - 1

	if trackCount < 5 {
		stop = trackCount
	}

	for i := 0; i <= stop; i++ {
		tempTicker.Tracks = append(tempTicker.Tracks, i)
	}

	tempTicker.TStop = tempTracks[tempTicker.Tracks[len(tempTicker.Tracks)-1]].Timestamp

	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Function run time
	tempTicker.Processing = time.Since(start) / 1000000 //Convert to ms
	//Return the struct as a json object.
	json.NewEncoder(w).Encode(tempTicker)
}

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

	start := time.Now()
	var tempTracks = getAllTracks()

	var index = -1
	for i := range tempTracks {
		if tempTracks[i].Timestamp > int64(newTimeStamp) {
			index = i
			break
		}
	}

	if index == -1 {
		//error 400?
		fmt.Println("No timestamp is bigger")
		return
	}

	var trackCount = len(tempTracks) - 1
	var tempTicker tickerStruct

	tempTicker.TLatest = tempTracks[trackCount].Timestamp
	tempTicker.TStart = tempTracks[index].Timestamp

	var stop = index + paging

	if index+paging > trackCount {
		stop = trackCount
	}

	for i := index; i <= stop; i++ {
		tempTicker.Tracks = append(tempTicker.Tracks, i)
	}

	tempTicker.TStop = tempTracks[tempTicker.Tracks[len(tempTicker.Tracks)-1]].Timestamp //Might be the correct behavior?
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Function run time
	tempTicker.Processing = time.Since(start) / 1000000 //Convert to ms
	//Return the struct as a json object.
	json.NewEncoder(w).Encode(tempTicker)
}
