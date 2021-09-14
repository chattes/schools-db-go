package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type School struct {
	SchoolId     int           `json:"id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	IsCatholic   bool          `json:"is_catholic"`
	Language     string        `json:"language"`
	Level        StringToSlice `json:"level"`
	City         string        `json:"city"`
	CitySlug     string        `json:"city_slug"`
	Board        string        `json:"board"`
	FraserRating float64       `json:"fraser_rating"`
	EQAORating   float64       `json:"eqao_rating"`
	Address      string        `json:"address"`
	Grades       string        `json:"grades"`
	Website      string        `json:"website"`
	PhoneNumber  string        `json:"phone_number"`
	Latitude     StringFloat64 `json:"latitude"`
	Longitude    StringFloat64 `json:"longitude"`
}

type AllSchools []School

func main() {
	fmt.Println("Starting DB Load...")

	asbPath, err := filepath.Abs("all-schools.json")
	if err != nil {
		fmt.Println("Error finding path")
	}

	jsonFile, err := os.Open(asbPath)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	schoolsBytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	var result AllSchools

	err = json.Unmarshal([]byte(schoolsBytes), &result)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All good...")

	TestMyFunc()

}
