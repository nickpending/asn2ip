package api

import (
	"asn2ip/util"
	"log"
)

func HandleASNQuery(asn string, ipv4Flag, ipv6Flag, debugFlag bool) {
	if ipv4Flag && ipv6Flag {
		util.PrintUsage()
		return
	}

	details, rawRequest, rawResponse, err := FetchASNDetails(asn, debugFlag)
	if err != nil {
		util.PrintError("ASN details", asn, err, debugFlag)
		return
	}

	if debugFlag {
		log.Printf("Request: %s\nResponse: %s\n", rawRequest, rawResponse)
	}

	if ipv4Flag {
		for _, prefix := range details.Data.IPv4Prefixes {
			util.PrintPrefix(prefix, debugFlag)
		}
	} else if ipv6Flag {
		for _, prefix := range details.Data.IPv6Prefixes {
			util.PrintPrefix(prefix, debugFlag)
		}
	} else {
		util.PrintUsage()
	}
}

func HandleIPQuery(ip string, debugFlag bool) {
	details, rawRequest, rawResponse, err := FetchIPDetails(ip, debugFlag)
	if err != nil {
		util.PrintError("IP details", ip, err, debugFlag)
		return
	}

	if debugFlag {
		log.Printf("Request: %s\nResponse: %s\n", rawRequest, rawResponse)
	}

	if len(details.Data.Prefixes) == 0 {
		util.PrintNoPrefixes(ip, debugFlag)
	} else {
		for _, prefix := range details.Data.Prefixes {
			asn := util.ASN{
				ASN:         prefix.ASN.ASN,
				Name:        prefix.ASN.Name,
				Description: prefix.ASN.Description,
				CountryCode: prefix.ASN.CountryCode,
			}
			util.PrintIPPrefix(prefix.PrefixDetails, asn, ip, debugFlag)
			break
		}
	}
}

func HandlePrefixQuery(prefix string, debugFlag bool) {
	details, rawRequest, rawResponse, err := FetchPrefixDetails(prefix, debugFlag)
	if err != nil {
		util.PrintError("prefix details", prefix, err, debugFlag)
		return
	}

	if debugFlag {
		log.Printf("Request: %s\nResponse: %s\n", rawRequest, rawResponse)
	}

	if len(details.Data.ASNs) == 0 {
		util.PrintNoASNs(prefix, debugFlag)
	} else {
		util.PrintPrefixInfo(details.Data.Prefix, details.Data.Name, details.Data.Description, details.Data.ASNs, debugFlag)
	}
}
