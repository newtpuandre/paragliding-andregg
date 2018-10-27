package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAdminTrackCount(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)
	defer clearTrackCol(t, testDB)

	var testTrack Track
	testTrack.Pilot = "test"

	insertTrack(&testTrack, testDB)

	server := httptest.NewServer(http.HandlerFunc(AdminTrackCount))

	//Get whole Track with id 0
	req, err := http.Get(server.URL + "/admin/api/tracks_count")
	if err != nil {
		t.Fatal(err)
	}

	// Check the status code is what we expect.
	if status := req.StatusCode; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//Convert response to bits and then to string
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}
	bodyString := string(bodyBytes)

	if !strings.Contains(bodyString, "1") {
		t.Fatal("Count was not one")
	}

}

func TestAdminTracksDelete(t *testing.T) {
	dbTestInit()
	testDB := setupDB(t)

	var testTrack Track
	testTrack.Pilot = "test"

	insertTrack(&testTrack, testDB)

	server := httptest.NewServer(http.HandlerFunc(AdminTracksDelete))

	req, err := http.NewRequest("DELETE", server.URL+"/admin/api/tracks", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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

	//Convert response to bits and then to string
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Could not convert to string")
	}
	bodyString := string(bodyBytes)

	if !strings.Contains(bodyString, "1") {
		t.Fatal("Deletion count was not one")
	}
}
