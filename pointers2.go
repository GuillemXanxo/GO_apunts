package main

import "fmt"

func foo(t *int) {
	j := 101
	t = &j //We give the address of j to t
}

func boo(t *int) {
	k := 102
	*t = k //Redirecting t to point to k
}

func car(t int) {
	k := 103
	t = k
}

func main() { // --> What does it print?
	i := 100
	fmt.Println(i) //i= 100

	foo(&i)
	fmt.Println(i) //i= 100 --> t in foo is a copy and it is not allowed to change the pointer

	boo(&i)
	fmt.Println(i) //i= 102 --> We have redirected the pointer to the actual value of k(102)

	car(i)
	fmt.Println(i) //i= 102 --> car is given a copy, but our original doesnt see any change (need to send pointer to affect the original)
}
