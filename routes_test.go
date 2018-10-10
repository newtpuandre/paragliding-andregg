//Tests for routes.go file
package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIInfoRoute(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/igcinfo/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(APIInfoRoute)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected1 := `"info": "Service for IGC tracks."`
	expected2 := `"version": "v1"`

	if strings.Contains(rr.Body.String(), expected1) && strings.Contains(rr.Body.String(), expected2) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected1+" "+expected2)
	}

}

func TestIgcIDPost(t *testing.T) {
	/*
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
	*/
}

func TestIgcIDAll(t *testing.T) {
	//Check if we return a array with info
	/*
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
	*/
}

func TestIgcID(t *testing.T) {
	//See if we get correct info in return.

}

func TestIgcIDField(t *testing.T) {
	//Check that info we get in return are correct
}
