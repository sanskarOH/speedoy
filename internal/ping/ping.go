package ping

import (
	"fmt"
	"net"
	"time"
)

func Run(host string, attempts int) error {

	ips, err := net.LookupIP(host)
	var lost int
	var sum int
	var avg int
	var ipv4 string
	var port = "9999"

	if err != nil {
		return err
	}

	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4 = ip.String()
			break

		}
	}

	if ipv4 == "" {
		return fmt.Errorf("no IPV4 address found")
	}
	address := net.JoinHostPort(ipv4, port)

	for i := 0; i < attempts; i++ {

		start := time.Now()
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		elapsed := time.Since(start).Milliseconds()

		if err != nil {
			fmt.Println("Connection Failed to the address: " + address)
			lost++
			continue
		}
		conn.Close()

		sum += int(elapsed)

		fmt.Printf("Reply from %s time= %d ms\n", ipv4, elapsed)
		time.Sleep(1 * time.Second)
	}

	recieved := attempts - lost
	if recieved > 0 {
		avg = sum / recieved
	}

	if recieved > 0 {
		fmt.Printf("Ping successful on %s Packets sent = %d Packets recieved = %d  Latency= %dms ", ipv4, attempts, recieved, avg)

	} else {
		fmt.Printf("Ping unsuccessful on %s Packets sent = %d Packets recieved = %d  Latency= %dms ", ipv4, attempts, recieved, avg)
		fmt.Printf("Host unreacheable on TCP port.")
	}

	return nil

}
