package main

import "sync"

func main() {
	counter := SafeCounter{Counter: 0, Mutex: sync.Mutex{}}
	var wg sync.WaitGroup
	for i:= 0; i<1000 ;i++   {
		wg.Add(1)
		go func() {
			counter.Mutex.Lock()
			defer wg.Done()
			defer counter.Mutex.Unlock()
			counter.Counter += 1
		}()
	}
	wg.Wait()
	println(counter.Counter)
}

type SafeCounter struct {
	Counter int
	Mutex   sync.Mutex
}