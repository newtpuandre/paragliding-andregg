package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestTickerLatest(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	server := httptest.NewServer(http.HandlerFunc(TickerLatest))
	defer server.Close()

	var testTrack Track

	testTrack.Timestamp = time.Now().Unix()
	testTrack.Pilot = "test"

	insertTrack(&testTrack, testDB)

	//Make a GET request to the server
	req, err := http.Get(server.URL + "/paragliding/api/ticker/latest")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Convert response to bits and then to string
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}
	bodyString := string(bodyBytes)

	stringTime := strconv.Itoa(int(testTrack.Timestamp))

	if !strings.Contains(bodyString, stringTime) {
		t.Fatal("Could not retrive latest timestamp")
	}

}

func TestTicker(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	server := httptest.NewServer(http.HandlerFunc(Ticker))
	defer server.Close()

	var testTrack Track

	testTrack.Timestamp = time.Now().Unix()
	testTrack.Pilot = "test"

	insertTrack(&testTrack, testDB)

	//Make a GET request to the server
	req, err := http.Get(server.URL + "/paragliding/api/ticker/")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response struct
	var testTicker tickerStruct
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testTicker)

	if decoderr != nil {
		t.Fatal(decoderr)
	}

	if testTicker.TLatest != testTrack.Timestamp || len(testTicker.Tracks) > 1 {
		t.Fatal("Ticker mismatch expected data")
	}

}

func TestTickerTimestamp(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	//Create a new mux router and add routes and create a server
	m := mux.NewRouter()
	addRoutes(m)
	server := httptest.NewServer(m)

	var testTrack Track

	testTrack.Timestamp = 1540567057
	testTrack.Pilot = "test"

	insertTrack(&testTrack, testDB)

	//Make a GET request to the server
	req, err := http.Get(server.URL + "/paragliding/api/ticker/1540567056")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response struct
	var testTicker tickerStruct
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testTicker)

	if decoderr != nil {
		t.Fatal(decoderr)
	}

	if testTicker.TLatest != testTrack.Timestamp || len(testTicker.Tracks) > 1 {
		t.Fatal("Ticker mismatch expected data")
	}
}
