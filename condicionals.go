package main

import (
	"fmt"
)

func main() {
  var edat int

  if edat >= 18 {
    fmt.Println("Pots accedir")
  } else if edat == 0 {
    fmt.Println("Sisplau, digues una edat vàlida")
  } else {
    fmt.Println("No pots passar")
  }
}
