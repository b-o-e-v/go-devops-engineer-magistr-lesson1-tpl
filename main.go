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

	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		res, err := getResourceUsage()

		if err != nil {
			fmt.Println(err)
			return
		}

		status, err := server.ParseStatus(res)

		if err != nil {
			fmt.Println(err)
			return
		}

		status.CheckLoadAverage()
		status.CheckMemoryUsage()
		status.CheckDiskSpace()
		status.CheckNetworkUsage()
	}
}
