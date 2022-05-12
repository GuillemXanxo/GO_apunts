package main

import (
	"fmt"
)

func main(){
  saluda() //funcion basica 
  saludaArgumentos("Guillem", "Perez") //funcion que recibe argumentos
  m:= max(3, 5) //funcion que recibe argumentos y tiene single return
  fmt.Println(m)
  max, min := maxMin(3, 5) //funcion con multiple retorno
  fmt.Println(max, min)
}

//Les funcions poden tenir varis returns. Es declaren fora de main

func saluda() {
  fmt.Println("Hola")
}

func saludaArgumentos(nombre, apellido string) {
  fmt.Printf("Hola %s %s\n", nombre, apellido)
}

func max(a, b int)int {
  if a > b{
    return a
  }
  return b
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
