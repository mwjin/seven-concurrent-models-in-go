package chopstick

import (
	"sync"
)

type Chopstick struct {
	id  int
	use *sync.Mutex
}

func NewChopstick(id int) *Chopstick {
	return &Chopstick{
		id, &sync.Mutex{},
	}
}

func (chopstick *Chopstick) Grab() {
	chopstick.use.Lock()
}

func (chopstick *Chopstick) Put() {
	chopstick.use.Unlock()
}

func (chopstick *Chopstick) GetId() int {
	return chopstick.id
}
