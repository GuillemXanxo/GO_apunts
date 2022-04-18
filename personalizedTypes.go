package main

import (
	"fmt"
)

//Creació de tipus personalitzat
type any int //hem creat un alias de tipus integer
var anyActual any

func main() {
  anyActual = 2022
  fmt.Printf("%T", anyActual)
}
