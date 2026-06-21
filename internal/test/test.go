package test

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Check() error {
	url := "https://proof.ovh.net/files/100Mb.dat"
	var nr int
	var buf []byte
	var actual int
	buf = make([]byte, 1024)
	resp, err := http.Get(url)

	if err != nil {

		return err

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("bad status: %s\n", resp.Status)
		return fmt.Errorf("Borken status code")
	}
	fmt.Printf("Connected to the server successfully\n")
	fmt.Printf("Status: %s\n", resp.Status)
	startTime := time.Now()

	for {

		//read a chunk
		nr, err = resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if nr == 0 {
			break
		}
		actual += nr

	}

	elapsedTime := time.Since(startTime).Seconds()

	fmt.Printf("---Download Statistics---\n")
	fmt.Printf("Bytes: %d", resp.ContentLength)
	fmt.Printf("\nRecieved Actual: %d\n", actual)
	fmt.Printf("Total Time: %.2fs", elapsedTime)
	fmt.Printf("\nDownload Speed: %.2fs", elapsedTime)

	return nil
}
