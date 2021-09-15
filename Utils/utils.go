package utils

import (
	"encoding/json"
	"strconv"
	"strings"
)

type StringToSlice []string
type StringFloat64 float64

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

func (ss *StringToSlice) UnmarshalJSON(b []byte) error {

	var e error

	if len(b) > 0 {
		strInput := string(b)
		strUnquote, _ := strconv.Unquote(strInput)
		if strUnquote == "null" {
			return nil
		}
		if strUnquote == "" {
			*ss = nil
		}

		levels := strings.Split(strUnquote, ",")

		for _, val := range levels {
			*ss = append(*ss, strings.TrimSpace(val))
		}

	} else {
		return nil
	}

	return e
}

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
