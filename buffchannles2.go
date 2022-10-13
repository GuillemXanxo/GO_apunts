package main

import (
	"fmt"
)

func genNumbers2(from, to int, ch chan int) {
	for j := from; j <= to; j++ {
		ch <- j
		fmt.Printf("Published: %d\n", j)
	}
}

func main() {
	ch := make(chan int) // unbuffered channel
	from, to := 1, 10
	go genNumbers2(from, to, ch)
	for j := from; j <= to; j++ {
		fmt.Printf("Received: %d\n", <-ch)
	}
}
