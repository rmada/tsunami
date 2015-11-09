package main

import (
	"crypto/tls"
	"io"
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
		client := &http.Client{}
		if fw.target.Scheme == "https" {
			//Skip certificate verify for performance
			secureTransport := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client = &http.Client{Transport: secureTransport}
		}
		for {
			if fw.dead {
				return
			}
			var body io.Reader;
			req,_ := http.NewRequest("GET", fw.target.String(), body)
			client.Do(req)
			fw.RequestCounter += 1 //Worker specific counter
			requestChan <- true
		}
	}()
}

func (fw *floodWorker) Kill() {
	fw.dead = true
	fw.exitChan <- fw.id
}
