package main

import (
	"fmt"
	"net"
	"projects/pkg/utils"
	"strings"
)

func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			addr, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addr {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func main() {
	utils.Print()
	ip := InternalIP()
	fmt.Println(ip)
}
