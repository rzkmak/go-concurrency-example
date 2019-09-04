package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	jobs := make(chan string, 10)
	done := make(chan bool, 10)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGSTOP)

	go func() {
		<-c
		log.Println("terminate service")
		os.Exit(0)
	}()


	go printer(jobs, done)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		jobs <- text
		fmt.Println("")
		<-done
	}

}

func printer(jobs <-chan string, done chan<-bool) {
	for job := range jobs {
		time.Sleep(time.Second)
		log.Println("printing " + job)
		done <- true
	}
}