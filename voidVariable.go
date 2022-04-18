package main

import (
	"fmt"
)

func main(){
  _, min := maxMin(3, 5)
  fmt.Println(min)
  /* Guardamos el segundo valor retornado en la variable min, para ello descartamos el primer valor retornado sin guardarlo en una variable.
  Si lo guardamos en una variable que luego no usamos, el compilador va a dar problemas, por eso lo descartamos con _ */
}

func maxMin(a, b int)(max int, min int) {
  if a > b{
    max = a
    min = b
  } else {
    max = b
    min = a
  }
  return
}
