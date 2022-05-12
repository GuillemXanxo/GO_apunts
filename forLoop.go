package main

import (
	"fmt"
)

func main() {
  d := 1

  //FOR WITH BREAK AND CONTINUE (seria el while de C)
  for{
    d++
    if d > 100 { //Condició de sortida quan d sigui 101 surt del bucle
      break
    }
    if d % 2 != 0 {
      continue //Torna a l'inici del bucle sense seguir amb les ordres
    }
    fmt.Println(d) //Només arriba aquí els casos de d és menor a 101 i parells (no satisfan condició anterior)
  }

  //FOR TRADICIONAL AMB LES CONDICIONS A ENUNCIAT
  for i := 1; i <= 100; i++{
    if i % 2 == 0{
      fmt.Println(i)
    }
  }

  // Exemple males practiques
  for i := 1; i <=100; i++ {
    if i%3 == 0 {
      if i%5 == 0{
        fmt.Println("FizzBuzz")
      } else {
        fmt.Println("Fizz")
      }
    } else if i%5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(i)
    }
  }

  // El mateix condicional en bones practiques
  for i := 1; i <=100; i++ {
    if i%3 == 0 && i%5 == 0{
      fmt.Println("FizzBuzz")
      continue
    }
    if i%3 == 0 {
      fmt.Println("Fizz")
      continue
    }
    if i%5 == 0 {
      fmt.Println("Buzz")
      continue
    }
    fmt.Println(i)
  }
}


