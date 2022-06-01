package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type Counter struct {
	count int
}

func NewCounter(count int) *Counter {
	return &Counter{count}
}

func (counter *Counter) increment() {
	mutex.Lock()
	defer mutex.Unlock()
	counter.count++
}

func (counter *Counter) getCount() int {
	return counter.count
}

func main() {
	counter := NewCounter(0)
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				counter.increment()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Expected: 20000, Result: %d\n", counter.getCount())
}
