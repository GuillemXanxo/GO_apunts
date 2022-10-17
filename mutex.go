package main

/* Mutex --> Mutual exclusion:
It keeps track of which routine is having acces to s variable
So action like read and write don't happen at the same time, that would cause a panic */

import (
	"fmt"
	"sync"
)

func main() {

	//Without Mutex
	m1 := map[int]int{}
	go writeLoop(m1) // --> this routine is trying to write while another routine is trying to read --> panic
	go readLoop(m1)

	//With Mutex
	m2 := map[int]int{}
	mux := &sync.Mutex{}
	go writeLoopMutex(m2, mux)
	go readLoopMutex(m2, mux)

	//With RWMutex
	m3 := map[int]int{}
	rwmux := &sync.RWMutex{} //--> A RWMutex let readers read concurrently but only allows one reader to access it, so we can have multiple reader at the same time
	go writeLoopRWMutex(m3, rwmux)
	go readLoopRWMutex("-", m3, rwmux)
	go readLoopRWMutex("&", m3, rwmux)
	go readLoopRWMutex("*", m3, rwmux)

	// stop program from exiting, must be killed
	block := make(chan struct{})
	<-block
}

func writeLoopRWMutex(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock() // --> we open the acces to that variable for this function, while this is open for this function it is closed to the rest.
			m[i] = i
			mux.Unlock() // --> we close the mutex so other function can have access to the variable
		}
	}
}

func readLoopRWMutex(s string, m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(k, s, v)
		}
		mux.RUnlock()
	}
}

func writeLoopMutex(m map[int]int, mux *sync.Mutex) {
	for {
		for i := 0; i < 100; i++ {
			mux.Lock() // --> we open the acces to that variable for this function, while this is open for this function it is closed to the rest.
			m[i] = i
			mux.Unlock() // --> we close the mutex so other function can have access to the variable
		}
	}
}

func readLoopMutex(m map[int]int, mux *sync.Mutex) {
	for {
		mux.Lock()
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
		mux.Unlock()
	}
}

func writeLoop(m map[int]int) {
	for {
		for i := 0; i < 100; i++ {
			m[i] = i
		}
	}
}

func readLoop(m map[int]int) {
	for {
		for k, v := range m {
			fmt.Println(k, "-", v)
		}
	}
}
