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
