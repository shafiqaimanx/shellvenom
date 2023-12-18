package core

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
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

func CheckIPInterfaceName(lhostFlag string) (bool, string) {
	var result string
	check := false

	ipList := GettingIPAddresses()
	for i:=0;i<len(ipList);i++ {
		nameListIndex := GettingIPInterfaceName(ipList[i])
		if nameListIndex != lhostFlag {
			check = false
		} else {
			result = IPToStringSlice(ipList)[i]
			check = true
			break
		}
	}
	return check, result
}

func CheckIPrange(lhostFlag string) bool {
	check := false

	ipPattern := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	ipRegex := regexp.MustCompile(ipPattern)
	if ipRegex.MatchString(lhostFlag) {
		check = true
	} else {
		check = false
	}
	return check
}

func CheckPortRange(lportFlag string) (bool, int) {
	port, err := strconv.Atoi(lportFlag)
	if err != nil || port > 65535 {
		return false, 0
	}
	return true, port
}

func CompareIPwithIntFaceAndRegex(checkIPIntFace, checkIPRangeRgx bool, lhostFlag *string) (check bool, choosenIFace string) {
	check = true
	if (!checkIPIntFace && !checkIPRangeRgx) {
		check = false
	} else if checkIPIntFace {
		{}
	} else if checkIPRangeRgx {
		choosenIFace = *lhostFlag
	}
	return check, choosenIFace
}