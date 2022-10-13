package main

import (
	"fmt"
)

func genNumbers1(from, to int, ch chan int) {
	for j := from; j <= to; j++ {
		ch <- j
		fmt.Printf("Published: %d\n", j)
	}
}

func main() {
	ch := make(chan int, 10) // buffered channel
	from, to := 1, 10
	go genNumbers1(from, to, ch)
	for j := from; j <= to; j++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}

//CHECK DIFFERENCES WITH buffchannels2.go
//a buffered channel waits without blocking the loop while the unbuff channel blocks it.
