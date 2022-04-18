package main

import (
	"fmt"
)

func main() {
 
  var x [5]int //Array de 5 indexes de integers --> fixa i no variables ni en tipus ni tamany == [0,0,0,0,0]
  y := [5]int{1, 1, 1, 1, 1} //Array declarat i inicialitzat
  fmt.Println(x)
  fmt.Println(y)

  x[1] = 1 //Canviem el valor de un dels index
  fmt.Println(x)

  /*
  RANGE (sintactic sugar for a for loop)
  el mètode range ens retorna dues variables i = index i v = valor
  */
  for i, v := range y {fmt.Println(i, " ", v)}
  // en cas que volguem anular index ho fem de la seguent manera per tal de recorrer array:
  for _, v := range y {fmt.Println(v)}

}
