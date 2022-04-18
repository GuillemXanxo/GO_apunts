package main

import (
	"fmt"
)

var(
  ciutat string = "Barcelona"
)

func main() {
  frase := fmt.Sprint("La ciutat de ", ciutat, " està a Catalunya") //amb les , també podem concatenar
  //mètode Sprint de fmt ens crea un sol string del resultant de concatenar en un sol string i guardar-ho en una variable. No imprimeix.

  fmt.Println(frase)

}
