package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestWebhookNewTrack(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	server := httptest.NewServer(http.HandlerFunc(WebhookNewTrack))

	//Struct that we pass with the request
	var hookPost webhookStructResponse
	hookPost.WebhookURL = "http://test.test/test"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(hookPost)

	req, err := http.Post(server.URL+"/paragliding/api/webhook/new_track", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}
	bodyString := string(bodyBytes)

	//Is result correct?
	if !strings.Contains(bodyString, "0") {
		t.Fatal("Webhook ID is not correctly set")
	}

}

func TestWebhookIDGet(t *testing.T) {
	//Reset some variables
	var emptyArray []int
	webhookID = emptyArray
	lastWebhookID = 0

	dbTestInit()
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	//Create a new mux router and add routes and create a server
	m := mux.NewRouter()
	addRoutes(m)
	server := httptest.NewServer(m)

	//Struct that we pass with the request
	var hookPost webhookStructResponse
	hookPost.WebhookURL = "http://test.test/test"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(hookPost)

	req, err := http.Post(server.URL+"/paragliding/api/webhook/new_track", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Get whole Track with id 0
	req, err = http.Get(server.URL + "/paragliding/api/webhook/new_track/0")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response Track which we compare with the "correct" one
	var responseHook webhookStructResponse
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&responseHook)
	if decoderr != nil {
		t.Fatal(decoderr)
	}

	if responseHook.WebhookURL != hookPost.WebhookURL ||
		responseHook.MinTriggerValue != 1 {
		t.Fatal("Response contains unexpected data")
	}

}

func TestWebhookIDDelete(t *testing.T) {
	//Reset some variables
	var emptyArray []int
	webhookID = emptyArray
	lastWebhookID = 0

	dbTestInit()
	testDB := setupDB(t)
	defer clearHookCol(t, testDB)

	//Create a new mux router and add routes and create a server
	m := mux.NewRouter()
	addRoutes(m)
	server := httptest.NewServer(m)

	//Struct that we pass with the request
	var hookPost webhookStructResponse
	hookPost.WebhookURL = "http://test.test/test"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(hookPost)

	req, err := http.Post(server.URL+"/paragliding/api/webhook/new_track", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Get whole Track with id 0
	requ, err := http.NewRequest("DELETE", server.URL+"/paragliding/api/webhook/new_track/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(requ)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Check the status code is what we expect.
	if status := res.StatusCode; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response Track which we compare with the "correct" one
	var responseHook webhookStructResponse
	decoder := json.NewDecoder(res.Body)
	decoderr := decoder.Decode(&responseHook)
	if decoderr != nil {
		t.Fatal(decoderr)
	}

	if responseHook.WebhookURL != hookPost.WebhookURL ||
		responseHook.MinTriggerValue != 1 {
		t.Fatal("Response contains unexpected data")
	}

}
