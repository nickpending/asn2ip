package util

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type PrefixDetails struct {
	Prefix      string `json:"prefix"`
	ASN         int    `json:"asn,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Country     string `json:"country,omitempty"`
}

type ASN struct {
	ASN         int    `json:"asn"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CountryCode string `json:"country_code"`
}

func FetchDetails(url string, target interface{}, debugFlag bool) (string, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return "", "", err
	}

	if debugFlag {
		fmt.Printf("Request: %s %s\n", req.Method, req.URL)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request to %s: %v", url, err)
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", "", err
	}

	if debugFlag {
		fmt.Printf("Response: %s\n", string(body))
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		log.Printf("Error unmarshalling JSON response: %v", err)
	}
	return fmt.Sprintf("%s %s", req.Method, req.URL), string(body), err
}

func PrintUsage() {
	flag.Usage()
}

func PrintError(detailsType, identifier string, err error, debugFlag bool) {
	if debugFlag {
		log.Printf("Error fetching %s for %s: %v", detailsType, identifier, err)
	}
	fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", detailsType, err)
}

func PrintNoPrefixes(ip string, debugFlag bool) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	message := fmt.Sprintf("%s|%s|No prefixes found\n", timestamp, ip)
	fmt.Fprint(os.Stdout, message)
}

func PrintNoASNs(prefix string, debugFlag bool) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	message := fmt.Sprintf("%s|%s|No ASNs found\n", timestamp, prefix)
	fmt.Fprint(os.Stdout, message)
}

func PrintPrefix(prefix PrefixDetails, debugFlag bool) {
	if debugFlag {
		message := fmt.Sprintf("%s|%s|%s|%s\n", prefix.Prefix, prefix.Name, prefix.Description, prefix.Country)
		fmt.Fprint(os.Stdout, message)
	} else {
		fmt.Fprintln(os.Stdout, prefix.Prefix)
	}
}

func PrintIPPrefix(prefix PrefixDetails, asn ASN, ip string, debugFlag bool) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	message := fmt.Sprintf("%s|%s|AS%d|%s|%s|%s|%s|%s|%s\n", timestamp, ip, asn.ASN, asn.Name, asn.Description, asn.CountryCode, prefix.Prefix, prefix.Name, prefix.Description)
	fmt.Fprint(os.Stdout, message)
}

func PrintPrefixInfo(prefix, name, description string, asns []ASN, debugFlag bool) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	for _, asn := range asns {
		message := fmt.Sprintf("%s|%s|AS%d|%s|%s|%s|%s|%s|%s\n", timestamp, prefix, asn.ASN, asn.Name, asn.Description, asn.CountryCode, prefix, name, description)
		fmt.Fprint(os.Stdout, message)
	}
}
