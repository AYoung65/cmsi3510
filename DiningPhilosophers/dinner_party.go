package main

import (
	"fmt"
	"sync"
)

const NUM_PHILOSOPHERS int = 5
const NUM_CHOPSTICKS int = 5
const EAT_TIMES int = 3
const NUM_EATING_PHILOSOPHER int = 2

type Host struct {
	requestChannel chan *Philosopher
	
	quitChannel chan int

	eatingPhilosophers map[int]bool

	mu sync.Mutex
}

func (h *Host) manage() {
	for {
		if len(h.requestChannel) == NUM_EATING_PHILOSOPHER {
			finished := <-h.requestChannel 
			currentlyEating := make([]int, 0, NUM_PHILOSOPHERS)
			for index, eating := range h.eatingPhilosophers {
				if eating {
					currentlyEating = append(currentlyEating, index)
				}
			}
			fmt.Printf("%v have been eating, clearing plates %d\n", currentlyEating, finished.ID)

			h.eatingPhilosophers[finished.ID] = false
		}

		select {
		case <-h.quitChannel:
			fmt.Println("party is over")
			return
		default:
		}
	}

}

func main() {
	var wg sync.WaitGroup
	requestChannel := make(chan *Philosopher, NUM_EATING_PHILOSOPHER)
	quitChannel := make(chan int, 1)
	host := Host{
		requestChannel: requestChannel,
		quitChannel: quitChannel,
		eatingPhilosophers: make(map[int]bool)
	}

	for i := range(

	philos := make([]*Philosopher, NUM_PHILOSOPHERS)

	for i := range(NUM_PHILOSOPHERS) {
		philos[i] = &Philosopher{
			ID: i + 1,
			Name: "",
			LeftChopStick: chopsticks[i],
			RightChopStick: chopsticks[(i+1)%5],
			Host: &host
		}
	}

	go host.manage()


	for i in philos {
		go philosopher[i].eat(&wg)
	} 


	wg.Wait()
	host.quitChannel <- 1

	<-host.quitChannel



}

