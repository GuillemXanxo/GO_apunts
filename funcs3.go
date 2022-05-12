package main

import (
	"fmt"
)

type functionOperators struct {
  Name string
  Surname string
  Age int
}

func myFunction (opts functionOperators) error{
  if Name != nil {
    fmt.Println("Hello ", Name)
  } else if Surname != nil {
    fmt.Println("Hfrom the family ", Surname)
  } else if Age != nil {
    fmt.Println("your age is ", Age)
  }
}

func main(){
  //FUNCIONS DE OPTIONAL PARAMETERS
  myFunction(functionOperators{Name: "Guillem", Age: 32})
}
