package netutil

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIPs() ([]string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var result = make([]string, 0, 10)
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addresses, _ := netInterfaces[i].Addrs()
			for _, address := range addresses {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						result = append(result, ipNet.IP.String())
					}
				}
			}
		}
	}
	return result, nil
}

func GetLocalIPsStr() (string, error) {
	if ips, err := GetLocalIPs(); err != nil {
		return "", err
	} else {
		return strings.Replace(strings.Trim(fmt.Sprint(ips), "[]"), " ", ",", -1), nil
	}
}
