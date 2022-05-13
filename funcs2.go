/* Aixo passa amb primitive values i structs, pero no amb maps ni slices ja que
internament ja son pointers */

package main

import (
	"fmt"
)

func main(){
  a:= 3
  fmt.Println("a = ", a)
  Incrementa(a)
  fmt.Println("a = ", a)
  /* La funcion Incrmeneta trabaja sobre una copia que no retorna, por lo que el valor de a en el main sigue siendo de 3 */

  // 1a SOLUCION: la func retorna el valor y lo guardamos en una variable b, pero a no ha incrementado
  b := Incrementareturn(a)
  fmt.Println("b = ", b)

  // 2a SOLUCION: a Incrementa le mandamos un puntero que referencie a la variable a, asi actua directamente sobre esa variable
  IncrementaPointer(&a)
  fmt.Println("a = ", a)
}

func Incrementa(a int){
  a++
}

func Incrementareturn(a int)int{
  a++
  return a
}

func IncrementaPointer(a *int){
  *a++
}
