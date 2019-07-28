package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id              int
	leftCS, rightCS *ChopS
	eatTimes        int
}

var wait sync.Mutex

func (p Philo) eat() {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		wait.Lock()
		if (p.eatTimes) < 3 {
			//fmt.Printf("Philosopher %d is grabbing left chopstick...\n", p.id)
			p.leftCS.Lock()
			//fmt.Printf("Philosopher %d is grabbing right chopstick...\n", p.id)
			p.rightCS.Lock()
			wait.Unlock()
			p.eatTimes = p.eatTimes + 1
			fmt.Printf("starting to eat %d %dtimes\n", p.id, p.eatTimes)
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("finishing eating %d\n", p.id)

			p.leftCS.Unlock()
			p.rightCS.Unlock()

		} else {
			wait.Unlock()
		}
	}
}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat()
	}
	wg.Wait()
}
