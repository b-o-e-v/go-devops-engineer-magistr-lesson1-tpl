package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println(t)
	}
}
