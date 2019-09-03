package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, message string) {
	pings<-message
}

func pong(pings <-chan string, pongs chan<- string) {
	pongs <-<- pings
}
