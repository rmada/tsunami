package main

import (
	"fmt"
	"time"
)

func MaxSecondsEnforcer() {
	time.Sleep(time.Duration(*maxSeconds) * time.Second)
	fmt.Printf("Max seconds reached (%d)\n", *maxSeconds)
	if(*maxSeconds > 0) {
		GracefulExit()
	}
}
