package main

import (
	"fmt"
	"time"
)

func Outputter() {
	lastOutputTime := time.Now()
	lastDisplayRequests := 0

	for {
		t := time.Now()
		requestDifference := requestCounter - lastDisplayRequests
		timeDifference := float64(t.UnixNano() - lastOutputTime.UnixNano()) / 1000000000
		if(timeDifference == 0) {
			timeDifference = 1
		}
		reqs := float64(requestDifference) / timeDifference
		fmt.Printf("%02d:%02d:%02d -> %.1f req/s | %d new reqs | %d total reqs\n", t.Hour(), t.Minute(), t.Second(), reqs, requestDifference, requestCounter)
		lastDisplayRequests = requestCounter
		lastOutputTime = t
		time.Sleep(time.Duration(*displayInterval) * time.Millisecond)
	}
}
