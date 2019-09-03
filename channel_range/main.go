package main

import "log"

func main() {
	msg := make(chan int, 3)
	msg <- 1
	msg <- 2
	msg <- 3
	close(msg)

	for val := range msg {
		log.Println(val)
	}
}
