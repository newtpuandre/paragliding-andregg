package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//stored in memory
var webhookID []int
var lastWebhookID int

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
	//webhookStructList = append(webhookStructList, hookStruct)

	//Add ID to array for used ids
	webhookID = append(webhookID, lastWebhookID)
	//Remember to count up used ids

	var newID = strconv.Itoa(lastWebhookID)

	lastWebhookID++
	//Specify content type
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	//Return the struct as a json object.
	json.NewEncoder(w).Encode(newID)
}

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

		json.NewEncoder(w).Encode(newResponse)
	} else {
		//Return bad request
		http.Error(w, "", 404) //404 Not found
	}
}

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
		json.NewEncoder(w).Encode(newResponse)
	} else {
		//Return bad request
		http.Error(w, "", 404) //404 Not found
	}
}

//Go through all webhooks and check if webhooks should be invoked
func invokeWebHook() {
	Hooks := getWebHooks(&Credentials)

	for i := range Hooks {
		Hooks[i].NewTracks++
		if Hooks[i].NewTracks%Hooks[i].MinTriggerValue == 0 {
			postWebHook(&Hooks[i])
		}
		updateWebhook(&Hooks[i], &Credentials)
	}

}

func postWebHook(w *webhookStruct) {
	fmt.Println("Invoking " + w.WebhookURL)
	//Post to slack or discord.
}
