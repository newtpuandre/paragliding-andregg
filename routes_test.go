//Tests for routes.go file
package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

var URL = "http://igcviewer-andregg.herokuapp.com" //Change to where the server are running

func APIInfoRouteTest(t *testing.T) {
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

	if testInfo.Version != "v1" || testInfo.Info != "Service for IGC tracks." {
		t.Error("Unkown version or service")
	}

}

func IgcIDPostTest(t *testing.T) {
	//Do a post request and check if we get a id in return
}

func IgcIDAllTest(t *testing.T) {
	//Check if we return a array with info
}

func IgcIDTest(t *testing.T) {
	//See if we get correct info in return.
}

func IgcIDField(t *testing.T) {
	//Check that info we get in return are correct
}
