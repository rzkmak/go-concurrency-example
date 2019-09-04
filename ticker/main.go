package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 2)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("ticker executed")
			case <-done:
				log.Println("ticker aborted")
			}
		}
	}()

	time.Sleep(time.Second * 10)
	ticker.Stop()
	done <- true

	log.Println("ticker ended")
}
