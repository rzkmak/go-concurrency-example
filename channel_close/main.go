package main

import (
	"fmt"
	"strconv"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received a new job with id " + strconv.Itoa(j))
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Println("send jobs in id " + strconv.Itoa(i))
	}

	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}
