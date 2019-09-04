package main

import (
	"log"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	counter := Counter{}

	for i:= 0; i< 1000; i++ {
		wg.Add(1)
		go func() {
			counter.increment(&wg)
		}()
	}

	wg.Wait()

	log.Println("result : " + strconv.Itoa(counter.Counter))
}

type Counter struct {
	sync.Mutex
	Counter int
}

func (c *Counter) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	c.Counter++
	c.Unlock()
}