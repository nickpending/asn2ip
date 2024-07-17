**asn2ip** is a command-line tool that retrieves and displays AS information.
# Features
- Uses [bgpview.io](https://bgpview.io) api
- Query by ASN for information
	- IPv4 and IPv6 prefixes.
- Identify ASN member ship
	- By IP
	- By Prefix
- Output
	- TEXT formatted similar to [asnmap](https://github.com/projectdiscovery/asnmap) 
	- STDOUT
## Why does this *seem* like asnmap?
We love [asnmap](https://github.com/projectdiscovery/asnmap), but we didn't want to use the free IPtoASN database that was being used. We've observed different results when compared with data from  [bgpview.io](https://bgpview.io) and  [bgp.he.net](https://bgp.he.net/) At some point, I'm sure we'll want to merge this option into asnmap, but for now this is where we landed.
## Installation

1. Clone the repository:
    ```git clone https://github.com/nickpending/asn2ip.git```
2. Build the project:
    ```go build```
## Usage

asn2ip -a ASNUMBER [-ipv4 | -ipv6] [-v] | -i IPADDRESS [-v] | -p CIDR [-v]
### Examples

- Query ASN and display IPv4 prefixes:
  `asn2ip -a AS15169 -ipv4`

- Query IP address:
  `asn2ip -i 8.8.8.8`

- Query CIDR prefix:
  `asn2ip -p 8.8.8.0/24`
## Flags

  -a, -a  ASN to query, e.g., AS6431
  -i, -i  IP address to query
  -ipv4, -i  Display only IPv4 prefixes
  -ipv6, -i  Display only IPv6 prefixes
  -p, -p  CIDR prefix to query
  -v, -v  Verbose output
  -version, -v  Print the version and exit
## Development

To contribute to the project:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Submit a pull request.
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
## Authors

- **Rudy R** - *Initial work* - [nickpending](https://github.com/nickpending)
---
Enjoy using `asn2ip`! If you have any questions or feedback, please feel free to contact us.
