package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getResourceUsage() error {
	res, err := http.Get("http://srv.msk01.gigacorp.local/_stats")

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return err
	}

	fmt.Println(body)

	return nil
}

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		getResourceUsage()
	}
}
