package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("start:", time.Now())

	mySleep(2 * time.Second)

	fmt.Println("end:", time.Now())
}
