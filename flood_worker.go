package main

import (
	"bytes"
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

			//Client logic inside loop for future dynamic tokens implementation
			body := []byte(tokenizedBody.String())
			req, _ := http.NewRequest(*method, tokenizedTarget.String(), bytes.NewBuffer(body))
			if *method == "POST" {
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			}

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
