// Lo mismo de pasar por valor y no referencia, que pasa con slices? slices son punteros.

package main

import (
	"fmt"
)

func main(){
  values := []int{1,2,3,4,5}
  IncrementaSlice(values) 
  fmt.Println(values) // Imprimeix [2, 3, 4, 5, 6]
}

func IncrementaSlice(values []int)[]int {
  for i := range values {
    values[i]++
  } 
  return values
}

