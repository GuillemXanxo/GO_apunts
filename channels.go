package main

import (
	"fmt"
)

//As soon as something get in chanA or chanB it will execute the function
func routine(chanA, chanB chan int) {
	for {
		select { //select works only for channels and executes the first channel that arrives
		case x := <-chanA:
			fmt.Printf("New int from channel A: %d\n", x)
		case y := <-chanB:
			fmt.Printf("New int from channel B: %d\n", y)
		}
	}
}

func main() {
	chanA := make(chan int)
	chanB := make(chan int)
	//We call the function and every time something gets in any of those channels the function will be called. It is non-blocking.
	//The way we do that is by "go" keyword, everytime a channel is feeded it will be consumed, but this function being opened doesnt mean it is blocking our code.
	go routine(chanA, chanB)

	chanA <- 42 //It will first receive info through chanA, routine executes case x. And so on. routine will be called 4 times in the order channels are feeded.
	chanB <- 24

	chanB <- 24
	chanA <- 42

}

//CONTINUE WITH waitgroups.go
