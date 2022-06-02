package philosopher

import (
	"fmt"
	"math/rand"
	"philosopher/chopstick"
)

type Philosopher struct {
	name    string
	first   *chopstick.Chopstick
	second  *chopstick.Chopstick
	randNum int
}

func NewPhilosopher(name string, left *chopstick.Chopstick, right *chopstick.Chopstick) *Philosopher {
	var first, second *chopstick.Chopstick
	if left.GetId() < right.GetId() {
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

func (philosopher *Philosopher) UseChopstick() {
	philosopher.first.Grab()
	fmt.Println(philosopher.name, "Grabs", philosopher.first.GetId())
	philosopher.second.Grab()
	fmt.Println(philosopher.name, "Grabs", philosopher.second.GetId())
	philosopher.second.Put()
	fmt.Println(philosopher.name, "Puts", philosopher.second.GetId())
	philosopher.first.Put()
	fmt.Println(philosopher.name, "Puts", philosopher.first.GetId())
}
