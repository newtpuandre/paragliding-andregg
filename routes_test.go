//Tests for routes.go file
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	var IGCURL Url
	IGCURL.Url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"
	var APIPostURL = "/igcinfo/api/igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)

	req, err := http.NewRequest("POST", APIPostURL, b)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IgcIDPost)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected1 := `{"id": 0}`

	if strings.Contains(rr.Body.String(), expected1) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected1)
	}

}

func TestIgcIDAll(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/igcinfo/api/igc", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IgcIDAll)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected1 := `"["`
	expected2 := `"]"`

	if strings.Contains(rr.Body.String(), expected1) && strings.Contains(rr.Body.String(), expected2) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected1+" "+expected2)
	}
}

func TestIgcID(t *testing.T) {
	rr := httptest.NewRecorder()

	var IGCURL Url
	IGCURL.Url = "http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc"

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(IGCURL)

	request, erro := http.NewRequest("POST", "igcinfo/api/igc", b)
	if erro != nil {
		t.Fatal(erro)
	}

	//post data
	handler2 := http.HandlerFunc(IgcIDPost)
	handler2.ServeHTTP(rr, request)
	fmt.Println(rr.Body.String())
	req, err := http.NewRequest("GET", "/igcinfo/api/igc/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.

	handler := http.HandlerFunc(IgcID)
	handler.ServeHTTP(rr, req)
	fmt.Println(rr.Body.String())
	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"H_date": "2016-02-19 00:00:00 +0000 UTC","pilot": "Miguel Angel Gordillo","glider": "RV8","glider_id": "EC-XLL","track_length": 425.95571656352956}`

	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}

func TestIgcIDField(t *testing.T) {
	req, err := http.NewRequest("GET", "/igcinfo/api/igc/0/track_length", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IgcField)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `425.95571656352956`

	if rr.Body.String() == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
