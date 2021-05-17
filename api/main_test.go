package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"cryptoCurrencies/Routes"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestDBCall(t *testing.T) {

	// Grab our router
	router := Routes.SetupRoutes()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/v1/service/db?fsyms=BTC&tsyms=USD")

	// Assert we encoded correctly,
	// the request gives a 200
	assertEqual(t, http.StatusOK, w.Code, "")

	// Convert the JSON response to a map
	var response map[string]map[string]map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Grab the value & whether or not it exists
	_, exists := response["RAW"]

	// Make some assertions on the correctness of the response.
	assertNil(t, err, "")
	assertTrue(t, exists, "")
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func assertNil(t *testing.T, a interface{}, message string) {
	if a == nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != nil", a)
	}
	t.Fatal(message)
}

func assertTrue(t *testing.T, value bool, message string) {
	if value {
		return
	}
	if len(message) == 0 {
		message = "Not True"
	}
	t.Fatal(message)
}
