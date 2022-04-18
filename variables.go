package main

import (
	"fmt"
)

var (
  g = "global" //variable global sense tipus pq infereix
  r string //variable declarada no inicialitzada, com que no pot inferir tipus li hem de dir nosaltres
)
const (
  pi = 3.14 //constant global i sense tipus
  cognom = "Xanxo"
)

func main() {
	nom := "Guillem" //infereix tipus
	fmt.Println("Hello, " + nom)
  nom = "Marc"
  fmt.Println("Hello, " + nom)
  fmt.Println(5 < 10)
}

