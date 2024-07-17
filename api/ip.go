package api

import (
	"fmt"
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

func FetchIPDetails(ip string) (*IPDetails, error) {
	url := fmt.Sprintf("https://api.bgpview.io/ip/%s", ip)
	var details IPDetails
	err := util.FetchDetails(url, &details)
	return &details, err
}

