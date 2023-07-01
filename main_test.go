// main_test.go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedHTML, err := ioutil.ReadFile("index.html")
	if err != nil {
		t.Fatal(err)
	}

	expected := string(expectedHTML)
	if !strings.Contains(recorder.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}

