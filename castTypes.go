package main

import (
	"fmt"
	"strconv"
)

func main() {

  //Exemple de valors per defectes de variables no inicialitzades, s'inicien automàticament amb valor 0
	var a int
  var b string
  var c float64
  var d bool

  fmt.Printf("variable a %T = %+v\n", a, a) //%T representa el tipus de la variable, en aquest cas ho estem aplicant a la var a
  fmt.Printf("variable b %T = %q\n", b, b)  //%q representa el valor de la variable al que se li explica
  fmt.Printf("variable c %T = %+v\n", c, c) //%+v mostra el valor de la variable sobre el que l'apliquem
  fmt.Printf("variable d %T = %+v\n", d /*la variable sobre la que s'aplica %T*/, d /*la variable sobre la que s'aplica %+v */)
  
  //CAST de tipus de més petit a més gran
  var index int8 = 15
  var bigIndex int32
  bigIndex = int32(index)

  fmt.Println(index)
  fmt.Println(bigIndex)

  //CAST de tipus de més gran a més petit --> ES POT PERDRE INFO
  var numerico int8
  var bigNumerico int32 = 64
  numerico = int8(bigNumerico)

  fmt.Println(numerico)
  fmt.Println(bigNumerico)

  //CAST de un int a un float
  var entero int8 = 32
  var flotante float64
  flotante = float64(entero)

  fmt.Println(entero)
  fmt.Printf("%.2f\n",flotante) //%.2f ens mostra fins a dos decimals de la variable sobre que apliquem

  //CAST de int a string (itoa)
  //importo package strconv

  nom := "Joan"
  edat:= 32

  fmt.Println("Hola " + nom + ", tens " + strconv.Itoa(edat) + " anys.")

  //CAST de string a int (atoi)

  
  

}
