package main

import (
  "fmt"
  "errors" //Importem package errors
)

func calcRemainder(numerator, denominator int) (int, error) { //la funcio retorna normal o lo ultim que retorna es error
  if denominator == 0 {
    return 0, errors.New("denominator is 0") //si es compleix condicio retornem un error
  }
  return numerator / denominator, nil //si tot va be, error que retornem es nil i a main controlem per nil
}

func main() {
  value1 := 300
  value2 := 10
  value3 := 0

  result1, err1 := calcRemainder(value1, value2)
  if err1 != nil { //Al retorn duna funcio sempre es fa el control de errors
    fmt.Println("An error occured in result1")
  } 
  fmt.Println(result1)
  

  result2, err2 := calcRemainder(value1, value3)
  if err2 != nil {
    fmt.Println("An error occured in result2")
  }
  fmt.Println(result2)
  

}
