package main

import (
	"log"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 2)
	log.Println("inside the timer")
	time.Sleep(time.Second * 2)
	<-timer.C
	log.Println("timer expired")

	secTimer := time.NewTimer(time.Second * 2)

	go func() {
		<-secTimer.C
		log.Println("secTimer expired")
	}()

	stop := secTimer.Stop()
	if stop {
		log.Println("secTimer stopped")
	}
}
