package main

import (
	"fmt"
	"net/http"
)

func floodWorker(id int, url string) {
	defer func() {
		threadExit <- true
		if *verbose {
			fmt.Printf("Thread %x ended\n", id)
		}
	}()
	if *verbose {
		fmt.Printf("Thread %x started\n", id)
	}
	for {
		_, _ = http.Get(url)
		totalRequestsSent += 1
	}
}
