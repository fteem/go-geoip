package response_parser

import (
	"encoding/json"
	"github.com/fteem/go-geoip/util"
)

type Response struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

func (r Response) String() string {
	output := ""

	output += "IP: " + r.IP

	if r.City != "" {
		output += "\n"
		output += "City: " + r.City + " (" + r.RegionCode + ")"
	}

	if r.CountryName != "" {
		output += "\n"
		output += "Country: " + r.CountryName
	}

	if r.CountryCode != "" {
		output += " (" + r.CountryCode + ")"
	}

	return output
}

func Parse(data []byte) Response {
	var response Response

	err := json.Unmarshal(data, &response)
	util.Check(err)

	return response
}
