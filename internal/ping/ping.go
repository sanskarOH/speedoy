package ping

import (
	"fmt"
	"net"
)

func Run(host string) error {

	ips, err := net.LookupIP(host)

	var ipv4 string

	if err != nil {
		return err
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4 = ip.String()
			break

		}
	}

	fmt.Println(ipv4)
	return nil

}
