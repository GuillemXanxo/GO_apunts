package main

import (
	"fmt"
	"math/rand"
)

func main() {

  // Primera manera --> la variable de condicio es declara fora en un scope mes gran que la codnicio
  var edat int

  if edat >= 18 {
    fmt.Println("Pots accedir")
  } else if edat == 0 {
    fmt.Println("Sisplau, digues una edat vàlida")
  } else {
    fmt.Println("No pots passar")
  }

  //Segona manera --> declarem la variable condici'o dins del scope de condicio

  if edat := rand.Intn(99); edat == 0 {
    fmt.Println("Encara no tens edat")
  } else if edat < 18 || edat >= 65 {
    fmt.Println("Toca treballar")
  } else {
    fmt.Println("Jubila't")
  }


}
