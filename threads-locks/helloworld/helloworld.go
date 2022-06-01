package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		fmt.Println("Hello, World!")
		wg.Done()
	}()
	wg.Wait()
}
