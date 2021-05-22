package ip

import (
	"net"
	"strings"
)

func Reverse(ip string) string {
	reversedIP := []string{}
	octets := strings.Split(ip, ".")

	for i := range octets {
		octet := octets[len(octets)-1-i]
		reversedIP = append(reversedIP, octet)
	}
	return strings.Join(reversedIP, ".")
}

func IsValidIPV4(ip string) bool {
	if len(ip) == 0 {
		return false
	}

	parsedIP := net.ParseIP(ip)
	if parsedIP == nil || parsedIP.To4() == nil {
		return false
	}

	return true
}
