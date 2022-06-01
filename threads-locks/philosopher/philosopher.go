package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type Chopstick struct {
	id  int
	use *sync.Mutex
}

func NewChopstick(id int) *Chopstick {
	return &Chopstick{
		id, &sync.Mutex{},
	}
}

func (chopstick *Chopstick) grab() {
	chopstick.use.Lock()
}

func (chopstick *Chopstick) put() {
	chopstick.use.Unlock()
}

func (chopstick *Chopstick) getId() int {
	return chopstick.id
}

type Philosopher struct {
	name    string
	first   *Chopstick
	second  *Chopstick
	randNum int
}

func NewPhilosopher(name string, left *Chopstick, right *Chopstick) *Philosopher {
	var first, second *Chopstick
	if left.getId() < right.getId() {
		first = left
		second = right
	} else {
		first = right
		second = left
	}
	return &Philosopher{
		name, first, second, rand.Intn(500),
	}
}

func (philosopher *Philosopher) useChopstick() {
	philosopher.first.grab()
	fmt.Println(philosopher.name, "grabs", philosopher.first.getId())
	philosopher.second.grab()
	fmt.Println(philosopher.name, "grabs", philosopher.second.getId())
	philosopher.second.put()
	fmt.Println(philosopher.name, "puts", philosopher.second.getId())
	philosopher.first.put()
	fmt.Println(philosopher.name, "puts", philosopher.first.getId())
}

func main() {
	chopsticks := [5]*Chopstick{}
	for i := 0; i < 5; i++ {
		chopsticks[i] = NewChopstick(i)
	}
	philosophers := [5]*Philosopher{}
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("Philosopher%d", i+1)
		left := chopsticks[i]
		var right *Chopstick
		if i+1 == len(chopsticks) {
			right = chopsticks[0]
		} else {
			right = chopsticks[i+1]
		}
		philosophers[i] = NewPhilosopher(name, left, right)
	}

	wg.Add(5)
	for _, philosopher := range philosophers {
		go func(p *Philosopher) {
			p.useChopstick()
			wg.Done()
		}(philosopher)
	}
	wg.Wait()
}
