package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var numOfPhil int = 5
var eatT int = 3
var numOfHost = 2
var host = make(chan bool, 2)

// var numEatingPhilo int = 2
type ChStick struct {
	sync.Mutex
}

type Philosopher struct {
	philNum                       int
	leftChopStick, rightChopStick *ChStick
}
type Host struct {
	permit chan int
}

func (p Philosopher) eat() {

	//Eating only 3 times
	count := 0
	for i := 0; i < eatT; i++ {

		host <- true
		p.leftChopStick.Lock()
		p.rightChopStick.Lock()
		fmt.Println("starting to eat", p.philNum+1)

		time.Sleep(2 * time.Second)

		p.leftChopStick.Unlock()
		p.rightChopStick.Unlock()
		fmt.Println("finishing eating", p.philNum+1)
		//time.Sleep(time.Second)
		count++
		<-host

	}
	wg.Done()
}

func main() {

	fmt.Println("Dining Philosopher problem ....")
	//host := new(Host)

	cSticks := make([]*ChStick, 5)
	for i := 0; i < numOfPhil; i++ {
		cSticks[i] = new(ChStick)
	}
	philos := make([]*Philosopher, 5)
	for i := 0; i < numOfPhil; i++ {
		philos[i] = &Philosopher{i, cSticks[i], cSticks[(i+1)%5]}

		wg.Add(1)
		// goroutine for the five philosophers
		go philos[i].eat()
	}
	wg.Wait()

}
