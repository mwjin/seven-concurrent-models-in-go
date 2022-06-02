package main

import (
	"fmt"
	"philosopher/chopstick"
	"philosopher/philosopher"
	"sync"
)

var wg sync.WaitGroup

func main() {
	chopsticks := [5]*chopstick.Chopstick{}
	for i := 0; i < 5; i++ {
		chopsticks[i] = chopstick.NewChopstick(i)
	}
	philosophers := [5]*philosopher.Philosopher{}
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("Philosopher%d", i+1)
		left := chopsticks[i]
		var right *chopstick.Chopstick
		if i+1 == len(chopsticks) {
			right = chopsticks[0]
		} else {
			right = chopsticks[i+1]
		}
		philosophers[i] = philosopher.NewPhilosopher(name, left, right)
	}

	wg.Add(5)
	for _, p := range philosophers {
		go func(p *philosopher.Philosopher) {
			p.UseChopstick()
			wg.Done()
		}(p)
	}
	wg.Wait()
}
