package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	name            int
}

func (p Philo) askHost(num chan int, askgroup *sync.WaitGroup) {
	fmt.Printf("Host was asked for permission to eat by philosopher %d \n", p.name)
	num <- 1
	fmt.Printf("Host allowed philosopher %d to eat. Philosophers eating: %d \n", p.name, len(num))
	askgroup.Done()
}

func (p Philo) tellHost(num chan int, tellgroup *sync.WaitGroup) {
	<-num
	tellgroup.Done()

}

func (p Philo) eat(num chan int) {
	for j := 0; j < 3; j++ {
		var askgroup sync.WaitGroup
		askgroup.Add(1)
		go p.askHost(num, &askgroup)
		askgroup.Wait()
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("starting to eat %d \n", p.name)
		time.Sleep(time.Second)
		fmt.Printf("finishing eating %d \n", p.name)
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		var tellgroup sync.WaitGroup
		tellgroup.Add(1)
		go p.tellHost(num, &tellgroup)
		tellgroup.Wait()
		waitgroup.Done()
	}
}

type Host struct {
	permissions int
}

func (h Host) allow() {

}

var waitgroup sync.WaitGroup

func main() {
	//CS initialization
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	waitgroup.Add(15)
	// Philos initialization
	philos := make([]*Philo, 5)
	num := make(chan int, 2)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
		go philos[i].eat(num)
	}

	waitgroup.Wait()

}
