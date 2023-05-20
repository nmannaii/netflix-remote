package utils

import (
	"net"
	"strings"
)

func GetLocalIpAddress() string {
	interfaces, _ := net.Interfaces()
	var wifiInterface []net.Addr

	for _, i := range interfaces {
		if i.Name == "Wi-Fi" {
			wifiInterface, _ = i.Addrs()
			break
		}
	}

	for _, addr := range wifiInterface {
		if strings.HasSuffix(addr.String(), "/24") {
			return addr.String()[0:strings.Index(addr.String(), "/")]
		}
	}

	return ""
}
