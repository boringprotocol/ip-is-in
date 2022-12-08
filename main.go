package main

import (
	"fmt"
	"net"
	"os"

	"github.com/vishvananda/netlink"
)

// A tool to print which private subnet to use based on what's detected in-use.
func main() {

	eth0, _ := netlink.LinkByName("eth0")

	addrs, _ := netlink.AddrList(eth0, netlink.FAMILY_V4)
	network := addrs[0].IPNet.String() // CIDR

	checkIPs := []string{
		"192.168.4.1",
		"10.1.4.1",
	}

	_, subnet, _ := net.ParseCIDR(network)
	for _, clientip := range checkIPs {
		ip := net.ParseIP(clientip)
		if !subnet.Contains(ip) {
			fmt.Printf("%s", clientip)
			os.Exit(0)
		}
	}

}
