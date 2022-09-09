package main

import (
	"fmt"
	"reflect"
)

/*
What is a variable in programming? if it was a struct it would look like:

type Variable struct {
	Name string
	Type string|int|float64  -->All primary types in Go
	Value The content we give to the variable Ex: "Hello world" or 42
	Address An address in the memory to know where it is stored
}

A pointer would look something like:

type Ptr struct {
	Name string
	Type string|int|float64 -->All primary types in Go. Sets to which type of variable can this pointer point
	Value The Variable.Address
}

*/
func main() {
	//Let's declare a variable:
	variable := 42
	fmt.Println("variable value is: ", variable)
	fmt.Println("variable type is: ", reflect.TypeOf(variable))
	fmt.Println("variable address is: ", &variable)

	//We declare a pointer to the variable
	pointer := &variable
	// & means address of, * means value of the variable that it is pointing to
	fmt.Println("pointer points to the value: ", *pointer, " -->dereferencing")
	fmt.Println("pointer value is: ", pointer, " it is the address of variable")

	variable2 := *pointer
	fmt.Println("variable2 has the value of the variable pointed by pointer (variable2 := *pointer), so we get:")
	fmt.Println("variable is: ", variable)
	fmt.Println("*pointer is: ", *pointer)
	fmt.Println("variable2 is: ", variable2)

	//What if we change value of *p?
	*pointer = 21
	fmt.Println("We set *p=21, so the value of the variable pointed by pointer is now 21:")
	fmt.Println("variable is: ", variable)
	fmt.Println("*pointer is: ", *pointer)
	fmt.Println("variable2 is: ", variable2, "this value was set before, so it remains as before")
}
