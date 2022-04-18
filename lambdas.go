package main

import (
	"fmt"
)

func main(){
  var generador func(int)int //generador es una referencia a una funcion que retorna int --> es una funcion lambda
  numero := 0
  generador = func(i int) int { //una lambda puede afectar variables fuera de su scope, como en este caso numero
    numero++
    return numero
  }
  fmt.Println(generador(numero), generador(numero), generador(numero))

  generador = func(i int) int { //la misma lambda puede asignarse a multiples funciones siempre que compartan signatura (tipos variable de entrada y de salida)
    return i * i
  }
  fmt.Println(generador(numero))
}

