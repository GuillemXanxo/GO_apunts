package main

import (
	"fmt"
)

func main(){
  energia := 78
  booster := false

  //Un switch que funciona com una cade de if, else if...cada case té la seva lògica
  switch {
  case energia == 100, booster == true:
    fmt.Println("Energia carregada")
  case energia < 100:
    energia += 20
    fmt.Println("Energia actual " , energia)
  default:
    fmt.Println("Error")
  }

  planta := 2

  switch planta {
  case 1, 3:
    fmt.Println("Articles llar")
  case 2:
    fmt.Println("Moda infantil")
  case 4:
    fmt.Println("Joguines")
  default:
    fmt.Println("Error")
  }

}

//switch no té pq tenir condició inicial, en aquest cas sempre entraria al switch
//en aquest cas la lògica estarà en cada case.
