package core

import (
	"fmt"
	"net"
)

func GettingIPAddresses() []net.IP {
	var ips []net.IP

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}

	for _, addr := range addresses {
		if ipnet, ok := addr.(*net.IPNet); ok || !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips
}

func GettingIPInterfaceName(IPAddress net.IP) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}

	for _, iface := range interfaces {
		if addrs, err := iface.Addrs(); err == nil {
			for _, addr := range addrs {
				if iip, _, err := net.ParseCIDR(addr.String()); err == nil {
					if iip.Equal(IPAddress) {
						return iface.Name
					}
				} else {
					continue
				}
			}
		} else {
			continue
		}
	}
	return ""
}

func IPToStringSlice(ips []net.IP) []string {
	ipStrings := make([]string, len(ips))
	for i, ip := range ips {
		ipStrings[i] = ip.String()
	}
	return ipStrings
}