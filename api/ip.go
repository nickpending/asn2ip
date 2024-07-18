package api

import (
	"asn2ip/util"
)

type IPDetails struct {
	Data struct {
		Prefixes []struct {
			util.PrefixDetails
			ASN struct {
				Name        string `json:"name"`
				CountryCode string `json:"country_code"`
				Description string `json:"description"`
				ASN         int    `json:"asn"`
			} `json:"asn"`
		} `json:"prefixes"`
	} `json:"data"`
}

func FetchIPDetails(ip string, debugFlag bool) (*IPDetails, string, string, error) {
	url := "https://api.bgpview.io/ip/" + ip
	var details IPDetails
	rawRequest, rawResponse, err := util.FetchDetails(url, &details, debugFlag)
	return &details, rawRequest, rawResponse, err
}
