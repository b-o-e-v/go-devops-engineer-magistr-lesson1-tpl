package poller

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Create(url string, retryCount int) func() chan string {
	return func() chan string {
		ch := make(chan string, 3)
		client := http.Client{Timeout: 5 * time.Second}
		errorCount := 0

		go func() {
			defer close(ch)

			for {
				if errorCount >= retryCount {
					fmt.Printf("Unable to fetch server statistic")
					break
				}

				response, err := client.Get(url)

				if err != nil {
					errorCount++
					fmt.Printf("failed to send request %s\n", err)
					continue
				}

				if response.StatusCode != http.StatusOK {
					errorCount++
					continue
				}

				body, err := io.ReadAll(response.Body)

				if err != nil {
					errorCount++
					fmt.Printf("failed to parse response %s\n", err)
					continue
				}

				response.Body.Close()
				ch <- string(body)
			}
		}()

		return ch
	}
}
