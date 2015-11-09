package main

import (
	"fmt"
	"time"
)

func MaxSecondsEnforcer() {
	time.Sleep(time.Duration(*maxSeconds) * time.Second)
	if *maxSeconds > 0 {
		fmt.Printf("Max seconds reached (%d)\n", *maxSeconds)
		GracefulExit()
	}
}
