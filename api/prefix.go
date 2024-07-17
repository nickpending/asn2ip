package api

import (
	"fmt"
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

func FetchPrefixDetails(prefix string) (*PrefixInfo, error) {
	url := fmt.Sprintf("https://api.bgpview.io/prefix/%s", prefix)
	var details PrefixInfo
	err := util.FetchDetails(url, &details)
	return &details, err
}

