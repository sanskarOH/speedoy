package ltc

import (
	"fmt"
	"net"
	"time"
)

func Check(host string, attempts int, port string, timeout int) error {

	ips, err := net.LookupIP(host)

	var lost int
	var sum int
	var avg int
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

	if ipv4 == "" {
		return fmt.Errorf("no IPV4 address found")
	}
	fmt.Printf("pinging %s [%s] on port %s\n", host, ipv4, port)
	// fmt.Printf("Resolved %s -> %s\n", host, ipv4)
	address := net.JoinHostPort(ipv4, port)

	for i := 0; i < attempts; i++ {

		start := time.Now()
		conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
		elapsed := time.Since(start).Milliseconds()

		if err != nil {
			fmt.Printf("Connection Failed: (%v)\n ", err)
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
		fmt.Printf("Ping unsuccessful on %s Packets sent = %d Packets recieved = %d  Latency= %dms\n", ipv4, attempts, recieved, avg)
		fmt.Printf("No response from service at TCP port %s .", port)
	}

	return nil

}
