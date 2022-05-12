package main

import (
	"fmt"
	"reflect"
)

func main() {
  x := []int{1, 2, 3, 4, 5}
  fmt.Println(x)

  //AFEGIR ELEMENTS A UN SLICE
  x = append(x, 6, 7, 8, 9, 10)
  fmt.Println(x)

  //Afegir un altre slice z=x+y
  y := []int{10, 9, 8, 7, 6, 5}
  z := []int{}
  z = append(x, y...)
  fmt.Println(z)

  //MANIPULACIÓ SLICE --> REVISAR EXERCICI DE SUBSLICES
  a := []int{}
  a = (z[9:11]) //z[9] inclòs, z[11] NO inclòs --> crees un slice de [z[9], z[10]]
  fmt.Println(a)

  //ELIMINAR ELEMENTS
  z = append(z[:9], z[11:]...) //Vull eliminar z[9] i z[10] per tant enganxo array de 0 a 8 i de 11 al final
  fmt.Println(z)

  //COPIAR SLICES
  original:= []int{1,2,3,4,5}
  copia := make([]int, 3, 3)
  n := copy(copia, original)
  fmt.Println(n, "numeros copiados: ", copia)
  /* la funcio copy copia un slice en otro hasta completar la capacidad de cualquiera de los dos. Retorna el numero de valores copiados, en este caso guardados en variable n */

  //Multidimensionals
  samarreta := []string{"XL", "Decathlon", "verd"}
  polo := []string{"M", "lacoste", "groc"}
  armari := [][]string{samarreta, polo}
  fmt.Println(armari)

  //Declarar un slice des de core Go:
  coreSlice := make([]int, 5, 100) //tipus de dades, dades inicials, capacitat màxima
  fmt.Println(coreSlice) //imprimirà el contingut del slice
  fmt.Println(len(coreSlice)) //imprimeix el número d'elements al slice LEN
  fmt.Println(cap(coreSlice)) //imprimeix la capacitat total del slice CAP

  /* Els slice són punters de memòria a un array. Cada vegada que canvia el length crea un array de la nova capacitat
per tant, formen un tipus en si mateix. Un slice no és comparable amb res que no sigui un slice.
Els lices es comparan amb la llibreria reflect, mètode DeepEqual. */

  s1 := []int{0, 1, 2}
  s2 := []int{0, 1, 2}
  if reflect.DeepEqual(s1, s2) {
    fmt.Println("Els slices son iguals!")
  } else {
    fmt.Println("Els slices no son iguals!")
  }


  //Proves
  //vull fer un array amb tot allò de slice prova que no sigui 0
  prova := []int{1,0,0,0,1,1}
  array1s := []int{}
  for i,v := range prova{
    if v != 0{
      array1s = append(array1s, prova[i])
    }
  } 
  fmt.Println(prova)
  fmt.Println(array1s)
}
