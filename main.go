package main

import (
	"asn2ip/api"
	"asn2ip/util"
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "1.0.0"

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
	verboseFlag := flag.Bool("v", false, "Verbose output")
	versionFlag := flag.Bool("version", false, "Print the version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "asn2ip collects ASN, IP, and prefix information from bgpinfo.io\n")
		fmt.Fprintf(os.Stdout, "Version: %s\n", version)
		fmt.Fprintf(os.Stdout, "Usage: -a ASNUMBER [-ipv4 | -ipv6] [-v] | -i IPADDRESS [-v] | -p CIDR [-v]\n")
		fmt.Fprintf(os.Stdout, "Example: -a AS6431 -ipv4 -v | -i 12.153.241.125 -v | -p 1.2.3.0/24 -v\n")
		fmt.Fprintf(os.Stdout, "\nFlags:\n")
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(os.Stdout, "  -%s, -%s  %s\n", f.Name, f.Name[:1], f.Usage)
		})
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
		api.HandleASNQuery(*asnFlag, *ipv4Flag, *ipv6Flag, *verboseFlag)
	} else if *ipFlag != "" {
		api.HandleIPQuery(*ipFlag, *verboseFlag)
	} else if *prefixFlag != "" {
		api.HandlePrefixQuery(*prefixFlag, *verboseFlag)
	} else {
		util.PrintUsage()
	}
}
