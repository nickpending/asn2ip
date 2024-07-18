package api

import (
	"asn2ip/util"
)

type ASNDetails struct {
	Data struct {
		IPv4Prefixes []util.PrefixDetails `json:"ipv4_prefixes"`
		IPv6Prefixes []util.PrefixDetails `json:"ipv6_prefixes"`
	} `json:"data"`
}

func FetchASNDetails(asn string, debugFlag bool) (*ASNDetails, string, string, error) {
	url := "https://api.bgpview.io/asn/" + asn + "/prefixes"
	var details ASNDetails
	rawRequest, rawResponse, err := util.FetchDetails(url, &details, debugFlag)
	return &details, rawRequest, rawResponse, err
}
