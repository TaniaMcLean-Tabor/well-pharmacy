package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

	for _, pharm := range pharmacies.Pharmacies {
		for i := 0; i < len(pharmacies.Pharmacies); i++ {
			if contains(pharm.Postcode) == false {
				pharmacies.Pharmacies = append(pharmacies.Pharmacies[:i], pharmacies.Pharmacies[i+1:]...)
			}
		}
	}

	json.NewEncoder(w).Encode(pharmacies)
	w.WriteHeader(http.StatusOK)
}

func contains(postcode string) bool {

	postcodes := []string{

		"M26",
		"M24",
		"M45",
		"M38",
		"M25",
		"M46",
		"M9",
		"M27",
		"M29",
		"M28",
		"M8",
		"M7",
		"M6",
		"M35",
		"M40",
		"M3",
		"M43",
		"M30",
		"M50",
		"M5",
		"M1",
		"M11",
		"M17",
		"M15",
		"M18",
		"M34",
		"M32",
		"M21",
		"M34",
		"M44",
		"M41",
		"M31",
		"M33",
		"M20",
		"M23",
		"M22",
		"M90",
		"M2",
	}

	postCodePrefix := postcode[:strings.IndexByte(postcode, ' ')]

	for _, p := range postcodes {
		if p == postCodePrefix {
			return true
		}
	}
	return false
}

// func getPostCodesFromImage() []string {

// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	client.SetLanguage("eng")
// 	client.SetImage("gm_g.png")
// 	text, _ := client.Text()
// 	fmt.Println(text)
// 	return strings.Split(text, "")
// }

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/pharmacies", returnAllPharmacies)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
