package api

import (
	"fmt"
	"asn2ip/util"
)

type ASNDetails struct {
	Data struct {
		IPv4Prefixes []util.PrefixDetails `json:"ipv4_prefixes"`
		IPv6Prefixes []util.PrefixDetails `json:"ipv6_prefixes"`
	} `json:"data"`
}

func FetchASNDetails(asn string) (*ASNDetails, error) {
	url := fmt.Sprintf("https://api.bgpview.io/asn/%s/prefixes", asn)
	var details ASNDetails
	err := util.FetchDetails(url, &details)
	return &details, err
}
