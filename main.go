package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Pharmacies struct {
	Pharmacies []Pharmacy `json:"data"`
}

type Pharmacy struct {
	Name         string `json:"name"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	City         string `json:"city"`
	Postcode     string `json:"postcode"`
	Phone        string `json:"phone"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllPharmacies(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Endpoint Hit: returnAllPharmacies")

	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened data.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var pharmacies Pharmacies

	json.Unmarshal(byteValue, &pharmacies)

	// for i := 0; i < len(pharmacies.Pharmacies); i++ {
	// 	fmt.Println("Name: " + pharmacies.Pharmacies[i].Name)
	// 	fmt.Println("AddressLine1: " + pharmacies.Pharmacies[i].AddressLine1)
	// 	fmt.Println("AddressLine2: " + pharmacies.Pharmacies[i].AddressLine2)
	// 	fmt.Println("City: " + pharmacies.Pharmacies[i].City)
	// 	fmt.Println("Postcode: " + pharmacies.Pharmacies[i].Postcode)
	// 	fmt.Println("Phone: " + pharmacies.Pharmacies[i].Phone)
	// }

	json.NewEncoder(w).Encode(pharmacies)

	w.WriteHeader(http.StatusOK)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pharmacies", returnAllPharmacies)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
