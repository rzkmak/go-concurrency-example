package main

import (
	"fmt"
	"strconv"
)

func main() {
	message := make(chan string)
	for i := 0; i<10; i++ {
		go func(i int) {
			message <- strconv.Itoa(i)
		}(i)
	}
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
