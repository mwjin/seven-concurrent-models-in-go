package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	answerReady := false
	answer := 0

	wg.Add(2)
	go func() {
		answer = 42
		answerReady = true
		wg.Done()
	}()
	go func() {
		for !answerReady {
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("The meaning of life is:", answer)
		wg.Done()
	}()
	wg.Wait()
}
