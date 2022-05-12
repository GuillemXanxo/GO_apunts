package main

import "fmt"

func main() {
  evenValues := []int{2, 4, 6, 8, 10, 12}

  //FOR-RANGE ITERATES OVER ARRAYS, SLICES , STRINGS AND MAPS

  for i, v := range evenValues {  // la variable i es index i v es el valor guardat en aquell index
    fmt.Println(i, v)
  }

  for _, v := range evenValues {  // si no volem index el podem anular amb _
    fmt.Println(v)
  }

  //FOR-RANGE ITERATION OVER A STRING THE V WILL BE THE UNICODE. --> SOLUTION:
  samples := []string{"hello", "you"}
  for _, sample := range samples {
    for i, unicode := range sample {
      fmt.Println(i, unicode, string(unicode))
    }
  }
}
