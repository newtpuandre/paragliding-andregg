//Tests for routes.go file
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var URL = "http://igcviewer-andregg.herokuapp.com" //Change to where the server are running

func TestAPIInfoRoute(t *testing.T) {
	//Check if the response contains correct fields and correct data
	var testInfo apiInfo

	var APIInfoURL = URL + "/igcinfo/api"

	res, err := http.Get(APIInfoURL)

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&testInfo)

	if err != nil {
		t.Error(err)
	}

	if testInfo.Version != "v1" || testInfo.Info != "Service for IGC tracks." {
		t.Error("Unkown version or service")
	}

}

func TestIgcIDPost(t *testing.T) {

	var IGCURL Url
	var IGCID Url_ID
	IGCURL.Url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"
	var APIPostURL = URL + "/igcinfo/api/igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)
	res, err := http.Post(APIPostURL, "application/json; charset=utf-8", b)

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&IGCID)

	if err != nil {
		t.Error(err)
	}

	if IGCID.Id < -1 {
		t.Error("ID is out of range.")
	}

}

func TestIgcIDAll(t *testing.T) {
	//Check if we return a array with info

	var APIPostURL = URL + "/igcinfo/api/igc"

	var testArray []int

	res, err := http.Get(APIPostURL)

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&testArray)

	if err != nil {
		t.Error(err)
	}

}

func TestIgcID(t *testing.T) {
	//See if we get correct info in return.

}

func TestIgcIDField(t *testing.T) {
	//Check that info we get in return are correct
}
