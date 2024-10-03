package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/b-o-e-v/go-devops-engineer-magistr-lesson1-tpl/server"
)

func getResourceUsage() (string, error) {
	response, err := http.Get("https://srv.msk01.gigacorp.local/_stats")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %s", response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	errorCount := 0

	for range ticker.C {
		res, err := getResourceUsage()
		if err != nil {
			errorCount++
		} else {
			status, err := server.ParseStatus(res)
			if err != nil {
				errorCount++
			} else {
				errorCount = 0

				status.CheckLoadAverage()
				status.CheckMemoryUsage()
				status.CheckDiskSpace()
				status.CheckNetworkUsage()
			}
		}

		if errorCount >= 3 {
			fmt.Println("Unable to fetch server statistic.")
			errorCount = 0
		}
	}
}
