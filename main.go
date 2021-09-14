package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type stringToSlice []string
type StringFloat64 float64

type School struct {
	SchoolId     int           `json:"id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`
	IsCatholic   bool          `json:"is_catholic"`
	Language     string        `json:"language"`
	Level        stringToSlice `json:"level"`
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

func (ss *StringFloat64) UnmarshalJSON(b []byte) error {
	var s interface{}

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch v := s.(type) {
	case float64:
		*ss = StringFloat64(v)
	case string:
		if v != "" {

			fl, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return err
			}
			*ss = StringFloat64(fl)

		} else {
			*ss = 0
		}
	}
	return nil
}

func (ss *stringToSlice) UnmarshalJSON(b []byte) error {

	var e error

	if len(b) > 0 {
		strInput := string(b)
		if strInput == "null" {
			return nil
		}
		if strInput == "" {
			*ss = nil
		}

		*ss = strings.Split(strInput, ",")

	} else {
		return nil
	}

	return e
}

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

}
