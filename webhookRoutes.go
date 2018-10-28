package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//stored in memory
var webhookID []int
var lastWebhookID int

//WebhookNewTrack parses incoming data and returns an id
func WebhookNewTrack(w http.ResponseWriter, r *http.Request) {

	//Decode incoming url
	var hookStruct webhookStruct
	hookStruct.WebhookURL = ""

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&hookStruct)
	if err != nil {
		//Malformed content body.
		http.Error(w, "Malformed content body", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	if hookStruct.WebhookURL == "" {
		//Malformed content body.
		http.Error(w, "Malformed content body", http.StatusBadRequest)

		return //Stop whatever we are doing..
	}

	if hookStruct.MinTriggerValue < 1 {
		hookStruct.MinTriggerValue = 1
	}

	hookStruct.WebhookID = lastWebhookID
	insertWebhook(&hookStruct, &Credentials)

	//Add ID to array for used ids
	webhookID = append(webhookID, lastWebhookID)

	var newID = strconv.Itoa(lastWebhookID)

	//Remember to count up used ids
	lastWebhookID++
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	err = json.NewEncoder(w).Encode(newID)
	if err != nil {
		fmt.Println(err)
	}
}

//WebhookIDGet gets a webhook which matches the provided id, if it exists.
func WebhookIDGet(w http.ResponseWriter, r *http.Request) {
	webhookStructList := getWebHooks(&Credentials)
	//Get parameters
	vars := mux.Vars(r)
	hookID := vars["webhook_id"]

	//Check if the parameter passed is an integer.
	i, err := strconv.Atoi(hookID)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "", 400) //400 bad request
		return
	}

	//Absolute value the integer. We dont accept negative numbers!
	if i < 0 {
		i = i * -1
	}

	//Could be better if i used a map
	var found = -1
	for j := range webhookStructList {
		if webhookStructList[j].WebhookID == i {
			found = j
		}
	}

	if found != -1 { //Is an int and not bigger than tracks in memory
		var newResponse webhookStructResponse
		newResponse.WebhookURL = webhookStructList[found].WebhookURL
		newResponse.MinTriggerValue = webhookStructList[found].MinTriggerValue
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		err = json.NewEncoder(w).Encode(newResponse)
	} else {
		//Return bad request
		http.Error(w, "", 400) //400 bad request
	}

	if err != nil {
		fmt.Println(err)
	}
}

//WebhookIDDelete deletes a webhook with a provided id, if it exists
func WebhookIDDelete(w http.ResponseWriter, r *http.Request) {
	webhookStructList := getWebHooks(&Credentials)
	//Get parameters
	vars := mux.Vars(r)
	hookID := vars["webhook_id"]

	//Check if the parameter passed is an integer.
	i, err := strconv.Atoi(hookID)

	if err != nil {
		fmt.Println(err)

		http.Error(w, "", 400) //400 bad request
		return
	}

	//Absolute value the integer. We dont accept negative numbers!
	if i < 0 {
		i = i * -1
	}

	//Could be better if i used a map
	var found = -1
	for j := range webhookStructList {
		if webhookStructList[j].WebhookID == i {
			found = j
		}
	}

	if found != -1 { //Is an int and not bigger than tracks in memory
		var newResponse webhookStructResponse
		newResponse.WebhookURL = webhookStructList[found].WebhookURL
		newResponse.MinTriggerValue = webhookStructList[found].MinTriggerValue

		deleteWebhook(&webhookStructList[found], &Credentials)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err = json.NewEncoder(w).Encode(newResponse)
	} else {
		//Return bad request
		http.Error(w, "", 404) //Bad Request
	}

	if err != nil {
		fmt.Println(err)
	}
}

//Go through all webhooks and check if webhooks should be invoked
func invokeWebHook(id int) {
	Hooks := getWebHooks(&Credentials)

	for i := range Hooks {
		Hooks[i].NewTracks++
		if Hooks[i].NewTracks%Hooks[i].MinTriggerValue == 0 {
			//Post and update webhook with last used id
			postWebHook(&Hooks[i], id)
			Hooks[i].LastTrackID = id
			updateWebhook(&Hooks[i], &Credentials)
		} else {
			//Just update the webhook without altering lastTrackID
			updateWebhook(&Hooks[i], &Credentials)
		}
	}

}

//postWebHook fills information and posts to the provided webhook.
func postWebHook(w *webhookStruct, id int) {
	fmt.Println("Invoking " + w.WebhookURL)
	//Struct that we pass with the request
	start := time.Now() //Function processing time

	tracks := getAllTracks(&Credentials)

	var single = -1    //Used if minTriggerValue == 1
	var multiple []int //Used if minTriggerValue > 1

	if w.MinTriggerValue == 1 { //Check minTriggerValue
		for i := range tracks { //Find id that we want to post
			if tracks[i].ID == id {
				single = i
			}
		}

	} else {
		//Add all new ids since last post
		for i := w.LastTrackID + 1; i <= w.LastTrackID+w.MinTriggerValue; i++ {
			multiple = append(multiple, i)
		}
	}

	var message discordMessage
	message.Content = "Latest timestamp: "

	//Find the latest timestamp and make it a string
	stringTimestamp := strconv.Itoa(int(tracks[len(tracks)-1].Timestamp))
	message.Content += stringTimestamp
	message.Content += ". New track ids are: ["

	if single == -1 { //minValueTrigger > 1

		for i := range multiple { //Range over values
			var trackID string

			if len(tracks) < 3 { //Fixing out of bounds error
				trackID = strconv.Itoa(tracks[multiple[i]-1].ID)
			} else {
				trackID = strconv.Itoa(tracks[multiple[i]].ID)
			}

			if i != len(multiple)-1 { //This is only for nice formatting
				message.Content = message.Content + trackID + ", "
			} else {
				message.Content = message.Content + "" + trackID
			}
		}
	} else { //Just add the one id we found
		trackID := strconv.Itoa(tracks[single].ID)
		message.Content += trackID
	}

	Processing := time.Since(start) //Time in MS
	message.Content += "]. Request took (" + Processing.String() + ")."

	hookURL := w.WebhookURL

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(message)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.Post(hookURL, "application/json", b)
	if err != nil {
		fmt.Println(err)
	}

	// Report if the post was successful or not
	if status := req.StatusCode; status != http.StatusNoContent {
		fmt.Printf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
