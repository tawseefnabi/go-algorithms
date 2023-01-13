// DNS records are mapping files that associate with DNS server whichever IP addresses each domain is associated
//
//	with, and they handle requests sent to each domain. The net package contains various methods to find general details of DNS records.
//
// Let's run some examples, to collect information about the DNS servers and the corresponding records of a target domain:
package main

import (
	"fmt"
	"net"
)

func main() {

	// The net.LookupIP() function accepts a string(domain-name) and
	// returns a slice of net.IP objects that contains host's IPv4 and IPv6 addresses.
	iprecords, _ := net.LookupIP("facebook.com")
	for key, ip := range iprecords {
		fmt.Println(key, "---->>>", ip)
	}

	// This is the abbreviation for canonical name. CNAMEs are essentially domain and subdomain text aliases to bind traffic. The net.LookupCNAME()
	// function accepts a host-name(m.facebook.com) as a string and returns a single canonical domain name for the given host.

	cname, _ := net.LookupCNAME("www.facebook.com")
	fmt.Println("cname", cname)

	var rect = map[string]int{}
	var rect2 = map[string]int{"height": 10}
	var rect3 map[string]int
	rect["height"] = 10
	fmt.Println(rect2, rect["height"], "rect3", rect3)
}
