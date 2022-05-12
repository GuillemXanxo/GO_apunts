package main

import (
	"fmt"
)

func addTo(base int, values ...int) []int {
  out := make([]int, 0, len(values))
  for _, v := range values {
    out = append(out, base+v)
  }
  return out
}

func main(){
  fmt.Println(addTo(3)) //no rep vlues per tant no pot aplicar res de la funcio
  fmt.Println(addTo(3, 2)) // torna array amb 2+3
  fmt.Println(addTo(3, 2, 4, 6, 8)) //torna array de [2+3, 4+3, 6+3, 8+3]

  a := []int{4, 3}
  fmt.Println(addTo(3, a...)) // a es un array [4, 3] la funcio retorna [4+3, 3+3]
  fmt.Println(addTo(3, []int{1, 2, 3, 4}...)) //el mateix que anterior pero li passem arra directament
}
