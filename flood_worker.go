package main

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type floodWorker struct {
	dead           bool
	exitChan       chan int
	id             int
	target         url.URL
	RequestCounter int
}

func (fw *floodWorker) Start() {
	go func() {
		defer fw.Kill()
		//Skip certificate verify for performance
		secureTransport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{}
		if fw.target.Scheme == "https" {
			client = &http.Client{Transport: secureTransport}
		}
		for {
			if fw.dead {
				return
			}
			client.Get(fw.target.String())
			fw.RequestCounter += 1 //Worker specific counter
			requestChan <- true
		}
	}()
}

func (fw *floodWorker) Kill() {
	fw.dead = true
	fw.exitChan <- fw.id
}
