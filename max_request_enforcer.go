package main

import (
	"fmt"
)

func MaxRequestEnforcer() {
	if *maxRequests < 0 {
		//Listen forever
		for {
			<-requestChan
			requestCounter++
		}
	} else {
		//Listen until max requests reached
		for requestCounter < *maxRequests {
			<-requestChan
			requestCounter++
		}
	}

	//Reached max requests
	fmt.Printf("Max requests reached (%d)\n", *maxRequests)
	GracefulExit()
}
