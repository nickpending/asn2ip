package api

import (
	"asn2ip/util"
)

func HandleASNQuery(asn string, ipv4Flag, ipv6Flag, verboseFlag bool) {
	if ipv4Flag && ipv6Flag {
		util.PrintUsage()
		return
	}

	details, err := FetchASNDetails(asn)
	if err != nil {
		util.PrintError("ASN details", asn, err, verboseFlag)
		return
	}

	if ipv4Flag {
		for _, prefix := range details.Data.IPv4Prefixes {
			util.PrintPrefix(prefix, verboseFlag)
		}
	} else if ipv6Flag {
		for _, prefix := range details.Data.IPv6Prefixes {
			util.PrintPrefix(prefix, verboseFlag)
		}
	} else {
		util.PrintUsage()
	}
}

func HandleIPQuery(ip string, verboseFlag bool) {
	details, err := FetchIPDetails(ip)
	if err != nil {
		util.PrintError("IP details", ip, err, verboseFlag)
		return
	}

	if len(details.Data.Prefixes) == 0 {
		util.PrintNoPrefixes(ip, verboseFlag)
	} else {
		for _, prefix := range details.Data.Prefixes {
			asn := util.ASN{
				ASN:         prefix.ASN.ASN,
				Name:        prefix.ASN.Name,
				Description: prefix.ASN.Description,
				CountryCode: prefix.ASN.CountryCode,
			}
			util.PrintIPPrefix(prefix.PrefixDetails, asn, ip, verboseFlag)
			break
		}
	}
}

func HandlePrefixQuery(prefix string, verboseFlag bool) {
	details, err := FetchPrefixDetails(prefix)
	if err != nil {
		util.PrintError("prefix details", prefix, err, verboseFlag)
		return
	}

	if len(details.Data.ASNs) == 0 {
		util.PrintNoASNs(prefix, verboseFlag)
	} else {
		util.PrintPrefixInfo(details.Data.Prefix, details.Data.Name, details.Data.Description, details.Data.ASNs, verboseFlag)
	}
}
