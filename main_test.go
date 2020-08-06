package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

func TestReturnAllPharmaciesHasJsonBody(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharmacies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllPharmacies)

	handler.ServeHTTP(rr, req)

	jsonPharmacies, err := ioutil.ReadAll(rr.Body)

	jsonPharmaciesMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(jsonPharmacies), &jsonPharmaciesMap)

	if err != nil {
		log.Fatalln(err)
	}

	if len(jsonPharmacies) <= 0 {
		t.Errorf("No response body returned")
	}
}

func TestReturnAllPharmaciesOnlyContainsManchesterPharmacies(t *testing.T) {
	req, err := http.NewRequest("GET", "/pharmacies", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllPharmacies)

	handler.ServeHTTP(rr, req)

	byteValue, _ := ioutil.ReadAll(rr.Body)

	var pharmacies Pharmacies

	json.Unmarshal(byteValue, &pharmacies)

	for _, pharm := range pharmacies.Pharmacies {
		if contains(pharm.Postcode) == false {
			t.Errorf("Only Manchester postcodes should be returned")
		}
	}
}

func contains(postcode string) bool {

	postcodes := []string{
		"M46",
		"M29",
		"M28",
		"Mao",
		"M26",
		"M45",
		"M25",
		"M27",
		"M44",
		"M31",
		"M24",
		"M40",
		"M35",
		"M18",
		"M34",
		"M33",
		"M21",
		"M20",
		"M23",
		"M22",
	}

	for _, p := range postcodes {
		if p == postcode {
			return true
		}
	}
	return false
}
