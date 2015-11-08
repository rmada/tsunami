package main

import (
	"fmt"
)

//If all workers die, exit
func WorkerOverseer() {
	for workerDeadCounter < *maxWorkers {
		deadId := <-exitChan
		if *verbose {
			fmt.Printf("Worker %x ended (Sent %d requests) \n", deadId, workers[deadId].RequestCounter)
		}
		workerDeadCounter += 1
	}
	fmt.Printf("All (%d) workers died", workerDeadCounter)
	GracefulExit()
}
