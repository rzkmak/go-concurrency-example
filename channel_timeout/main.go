package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go func() {
		time.Sleep(time.Second * 3)
		message <- "result"
	}()

	select {
	case msg := <-message:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}
}
