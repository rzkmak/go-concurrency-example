package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
}

func worker(done chan<- bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 1)
	fmt.Println("done")
	done <- true
}
