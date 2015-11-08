package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

//Command line args
var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	threads = kingpin.Flag("threads", "Amount of concurrent attacking threads.").Default("8").Int()
	url     = kingpin.Arg("url", "Target URL e.g http://google.com").Required().String()
)

//Attack stats variables
var (
	totalRequestsSent int
	threadCounter     int
	threadDoneCounter int
	threadExit        chan bool
)

func main() {
	kingpin.Parse()

	//Reflect arguments
	if *verbose {
		fmt.Printf("URL => %s\n", *url)
		fmt.Printf("Threads => %d\n", *threads)
	}

	threadExit := make(chan bool)

	//Start flood workers
	for threadCounter < *threads {
		go floodWorker(threadCounter, *url)
		threadCounter += 1
	}

	//Wait for workers to finish before exit
	for threadDoneCounter < *threads {
		<-threadExit
	}
	return
}
