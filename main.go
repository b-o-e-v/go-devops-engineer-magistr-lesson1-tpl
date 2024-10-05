package main

import (
	"github.com/b-o-e-v/go-devops-engineer-magistr-lesson1-tpl/poller"
	"github.com/b-o-e-v/go-devops-engineer-magistr-lesson1-tpl/profiler"
)

const (
	serverURL  = "http://srv.msk01.gigacorp.local/_stats"
	retryCount = 3
)

func main() {
	ch := poller.Create(serverURL, retryCount)

	for res := range ch() {
		stats, err := profiler.Parse(res)

		if err != nil {
			continue
		}

		stats.CheckLoadAverage()
		stats.CheckMemoryUsage()
		stats.CheckDiskSpace()
		stats.CheckNetworkUsage()
	}
}
