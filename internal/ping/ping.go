package ping

import (
	"fmt"
	"net"
)

func Run(host string) error {

	ips, err := net.LookupIP(host)

	if err != nil {
		return err
	}

	fmt.Println(ips)
	return nil

}
