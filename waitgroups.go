package main

import (
	"fmt"
	"sync"
)

//COMING FROM channels.go

func routine2(chanA, chanB chan int, wg *sync.WaitGroup) {
	for {
		select {
		case x := <-chanA:
			fmt.Printf("New int from channel A: %d\n", x)
		case y := <-chanB:
			fmt.Printf("New int from channel B: %d\n", y)
			wg.Done()
		}
	}
}

func main() {
	chanA := make(chan int)
	chanB := make(chan int)

	waitgroup := &sync.WaitGroup{}
	waitgroup.Add(2)
	go routine2(chanA, chanB, waitgroup)

	chanA <- 42
	chanB <- 24

	chanB <- 24
	chanA <- 42

	waitgroup.Wait() //This will wait until we have collected as many Done (line 17) as we have set in Add (line 27)

}
