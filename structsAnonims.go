package main

import (
	"fmt"
)

func main(){
  persona1 := struct {
    nom,cognom string
    edat int
  }{
    nom: "Maria",
    cognom: "Pérez",
    edat: 25,
  }

  fmt.Println(persona1)

}

/* Com que és un struct anònim l'hem de definir i poblar a l'hora de declarar-lo */
