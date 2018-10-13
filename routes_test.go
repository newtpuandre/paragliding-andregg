//Tests for routes.go file
//Improvements to be done: Clearer variable names.
package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestAPIInfoRoute(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	server := httptest.NewServer(http.HandlerFunc(APIInfoRoute))
	defer server.Close()

	//Make a GET request to the server
	req, err := http.Get(server.URL + "/igcinfo/api")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Create a struct to hold info
	var testInfo apiInfo
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testInfo)

	if decoderr != nil {
		t.Fatal(decoderr)
	}

	// Check the response body is what we expect. Dont check uptime. Is variable
	expected1 := "Service for IGC tracks."
	expected2 := `"version": "v1"`

	//Is info as expected?
	if testInfo.Info == expected1 && testInfo.Version == expected2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			testInfo.Info+" "+testInfo.Version, expected1+" "+expected2)
	}

}

func TestIgcIDPost(t *testing.T) {
	//Create new test server
	server := httptest.NewServer(http.HandlerFunc(IgcIDPost))
	defer server.Close()

	//Struct that we pass with the request
	var IGCURL URL
	IGCURL.URL = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)

	req, err := http.Post(server.URL+"/igcinfo/api/igc", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response struct
	var testID URLID
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testID)

	if decoderr != nil {
		t.Fatal(decoderr)
	}

	//Is result correct?
	if testID.ID < 0 {
		t.Fatal("ID not correct or correctly added!")
	}

}

func TestIgcIDAll(t *testing.T) {
	//Possible improvements. Better check for the API response

	//New test server
	server := httptest.NewServer(http.HandlerFunc(IgcIDAll))
	defer server.Close()

	//Get Request
	res, err := http.Get(server.URL + "/igcinfo/api/igc")
	if err != nil {
		t.Fatal(err)
	}

	//Convert response to bits and then to string
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}
	bodyString := string(bodyBytes)

	// Check the status code is what we expect.
	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Does the string contain a opening and closing bracket?
	//Aka is the response an array? Should check for an actual array?
	if !strings.Contains(bodyString, "[") {
		t.Fatal("Not an array in response.")
	}

	if !strings.Contains(bodyString, "]") {
		t.Fatal("Not an array in response.")
	}
}

func TestIgcID(t *testing.T) {
	//Create a new mux router and add routes and create a server
	m := mux.NewRouter()
	addRoutes(m)
	server := httptest.NewServer(m)

	//URL Struct we send as body
	var IGCURL URL
	IGCURL.URL = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)

	//Post IGCURL as json struct
	req, err := http.Post(server.URL+"/igcinfo/api/igc", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response struct
	var testID URLID
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testID)
	if decoderr != nil {
		t.Fatal(decoderr)
	}

	//Is the ID set and above 0?
	if testID.ID < 0 {
		t.Fatal("ID not correct or correctly added!")
	}

	//Get whole Track with id 0
	req, err = http.Get(server.URL + "/igcinfo/api/igc/0")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response Track which we compare with the "correct" one
	var testTrack Track
	decoder = json.NewDecoder(req.Body)
	decoderr = decoder.Decode(&testTrack)
	if decoderr != nil {
		t.Fatal(decoderr)
	}

	//Fill inn correct information and test the response againt it.
	var correctTrack Track
	correctTrack.H_date = "2016-02-19 00:00:00 +0000 UTC"
	correctTrack.Pilot = "Miguel Angel Gordillo"
	correctTrack.Glider = "RV8"
	correctTrack.Glider_id = "EC-XLL"
	correctTrack.Track_length = 443.2573603705269

	if testTrack != correctTrack {
		t.Fatal("Track info is different from test info")
	}

}

func TestIgcIDField(t *testing.T) {
	//Possible improvements: Check more than one field.

	//Create a new mux router and add routes and create a server
	m := mux.NewRouter()
	addRoutes(m)
	server := httptest.NewServer(m)

	//URL Struct we send as body
	var IGCURL URL
	IGCURL.URL = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)

	//Post IGCURL as json struct
	req, err := http.Post(server.URL+"/igcinfo/api/igc", "application/json", b)
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Response struct
	var testID URLID
	decoder := json.NewDecoder(req.Body)
	decoderr := decoder.Decode(&testID)
	if decoderr != nil {
		t.Fatal(decoderr)
	}

	//Is the ID set and above 0?
	if testID.ID < 0 {
		t.Fatal("ID not correct or correctly added!")
	}

	//Get whole Track with id 0 and pilot field
	req, err = http.Get(server.URL + "/igcinfo/api/igc/0/pilot")
	if err != nil {
		t.Fatal(err)
	}
	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Get response body and check against the "correct" pilot
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}

	var correctTrackInfo Track
	correctTrackInfo.Pilot = "Miguel Angel Gordillo"

	if strings.Contains(correctTrackInfo.Pilot, string(bodyBytes)) {
		t.Fatal("Track info is different from test info")
	}

}
