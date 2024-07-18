package main

import (
	"asn2ip/api"
	"asn2ip/util"
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "1.0.1"

func main() {
	// Initialize logging
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	asnFlag := flag.String("a", "", "ASN to query, e.g., AS6431")
	ipv4Flag := flag.Bool("ipv4", false, "Display only IPv4 prefixes")
	ipv6Flag := flag.Bool("ipv6", false, "Display only IPv6 prefixes")
	ipFlag := flag.String("i", "", "IP address to query")
	prefixFlag := flag.String("p", "", "CIDR prefix to query")
	debugFlag := flag.Bool("debug", false, "Print debug information including raw requests and responses")
	versionFlag := flag.Bool("version", false, "Print the version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "asn2ip collects ASN, IP, and prefix information from bgpview.io\n")
		fmt.Fprintf(os.Stdout, "Version: %s\n\n", version)
		fmt.Fprintf(os.Stdout, "Usage: -a ASNUMBER [-ipv4 | -ipv6] [-debug] | -i IPADDRESS [-debug] | -p CIDR [-debug]\n")
		fmt.Fprintf(os.Stdout, "Example: -a AS6431 -ipv4 -debug | -i 12.153.241.125 -debug | -p 1.2.3.0/24 -debug\n\n")
		fmt.Fprintf(os.Stdout, "Flags:\n")
		fmt.Fprintf(os.Stdout, "  -a, -a\tASN to query, e.g., AS6431\n")
		fmt.Fprintf(os.Stdout, "  -i, -i\tIP address to query\n")
		fmt.Fprintf(os.Stdout, "  -ipv4, -i\tDisplay only IPv4 prefixes\n")
		fmt.Fprintf(os.Stdout, "  -ipv6, -i\tDisplay only IPv6 prefixes\n")
		fmt.Fprintf(os.Stdout, "  -p, -p\tCIDR prefix to query\n")
		fmt.Fprintf(os.Stdout, "  -debug, -d\tPrint debug information including raw requests and responses\n")
		fmt.Fprintf(os.Stdout, "  -version, -v\tPrint the version and exit\n")
	}

	flag.Parse()

	if *versionFlag {
		fmt.Printf("asn2ip version %s\n", version)
		return
	}

	if *asnFlag != "" {
		if !*ipv4Flag && !*ipv6Flag {
			fmt.Fprintln(os.Stderr, "Error: -a option requires either -ipv4 or -ipv6 option.")
			flag.Usage()
			return
		}
		api.HandleASNQuery(*asnFlag, *ipv4Flag, *ipv6Flag, *debugFlag)
	} else if *ipFlag != "" {
		api.HandleIPQuery(*ipFlag, *debugFlag)
	} else if *prefixFlag != "" {
		api.HandlePrefixQuery(*prefixFlag, *debugFlag)
	} else {
		util.PrintUsage()
	}
}
