package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnAllPharmaciesHasStatusCodeOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharmacies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllPharmacies)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestReturnAllPharmaciesFiltersByManchesterPostCodes(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharmacies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllPharmacies)

	handler.ServeHTTP(rr, req)

}
