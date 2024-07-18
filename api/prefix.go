package api

import (
	"asn2ip/util"
)

type PrefixInfo struct {
	Data struct {
		Prefix      string     `json:"prefix"`
		IP          string     `json:"ip"`
		CIDR        int        `json:"cidr"`
		ASNs        []util.ASN `json:"asns"`
		Name        string     `json:"name"`
		Description string     `json:"description_short"`
	} `json:"data"`
}

func FetchPrefixDetails(prefix string, debugFlag bool) (*PrefixInfo, string, string, error) {
	url := "https://api.bgpview.io/prefix/" + prefix
	var details PrefixInfo
	rawRequest, rawResponse, err := util.FetchDetails(url, &details, debugFlag)
	return &details, rawRequest, rawResponse, err
}
