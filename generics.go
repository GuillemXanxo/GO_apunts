package main

import (
	"fmt"
)

func printAllInts(numbers []int) { //This function is limited to a list of int
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func printAll[T any](things []T) { //This function accepts T as a parameter, what is T? T is type any --> accepts everything
	for _, element := range things {
		fmt.Println(element)
	}
}

//What if we want to accept just numbers and nothing else? floats, ints, etc?
//We create a new type that includes all the types we want to consider:
type Number interface {
	int | float64
}

func printAllNumbers[T Number](things []T) { // This function only accepts T, which is int or float64
	for _, element := range things {
		fmt.Println(element)
	}
}

//What about returning functions? Let's create a divide function that accepts numbers and returns the result:

func divideNumbers[T Number](a, b T) T {
	return a / b
}

//GENERICS WITH CUSTOM TYPES
type CustomFloat float64

func divideFloats[T ~float64](a, b T) T { // ~ relaxes float64 to accept float64 and all custom types that resembles float64
	return a / b
}

//TYPE COMPARABLE
func contains[T comparable](elements []T, e T) bool { // --> this function accepts all types, but elements and e should be of the same type
	for _, s := range elements {
		if e == s {
			return true
		}
	}
	return false
}

func main() {
	//Custom types
	var z, x CustomFloat
	z = 3.14
	x = 5.63
	y := divideFloats(z, x) // --> z and x are not float64m their type is CustomFloat, but the function accepts them
	fmt.Println("y is: ", y)

	//Comparable
	list := []CustomFloat{3.12, 5.90, 5.63, 6}
	b := contains(list, x)
	fmt.Println("Result of contains is: ", b)

}
