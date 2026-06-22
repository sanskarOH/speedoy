package test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func selectServer(size string) (string, error) {
	serverList := map[string]string{
		"1mb":   "https://proof.ovh.net/files/1Mb.dat",
		"10mb":  "https://proof.ovh.net/files/10Mb.dat",
		"100mb": "https://proof.ovh.net/files/100Mb.dat",
		"1gb":   "https://proof.ovh.net/files/1Gb.dat",
	}
	if serverList[size] == "" {
		return "", errors.New("No server found")
	}
	return serverList[size], nil
}

func Check(size string) error {
	url, err := selectServer(size)
	if err != nil {
		return fmt.Errorf("Server was not found")
	}
	var nr int
	var buf []byte
	var actual int
	var milestone int
	buf = make([]byte, 64*1024)
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
	milestone = 25
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
		progress := (actual * 100 / int(resp.ContentLength))

		if progress >= milestone {
			fmt.Printf("\nProgress: %d percent", milestone)
			milestone += 25
		}

	}

	elapsedTime := time.Since(startTime).Seconds()

	speed := float64(actual*8) / elapsedTime / 1000000

	fmt.Printf("\n\n---Download Statistics---\n")

	fmt.Printf("File size: %.2f mb", (float64(resp.ContentLength) / (1024 * 1024)))
	fmt.Printf("\nFile Downloaded: %.2f mb\n", (float64(actual) / (1024 * 1024)))
	fmt.Printf("Total Time taken: %.2f s", elapsedTime)
	fmt.Printf("\nDownload Speed: %.2f Mbps", speed)

	return nil
}
