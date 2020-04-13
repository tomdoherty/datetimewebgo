package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDateHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/date", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	http.Handle("/date", DateHandler("hostname"))
	http.DefaultServeMux.ServeHTTP(rr, req)

	var result struct {
		Hostname string
		Date     string
	}

	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Errorf("failed to decode: %w", err)
	}

	if result.Hostname != "hostname" {
		t.Errorf("hostname not received")
	}
}
