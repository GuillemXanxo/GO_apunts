package main

import (
	"fmt"
)

//Els camps han de ser únics! Es defineixen fora de la funció main()
//Els noms en minuscula només seran accessibles des d'aquest package, si els volem fer publics han de ser en majúscules
type persona struct{
  nom string
  edat int
}

type enginyer struct{
  persona //promoció de persona, no herència
  realitzaPlanols bool
}
//persona passa a ser un tipus de dada

func main(){
  persona1 := persona{
    nom: "Guillem",
    edat: 32,
  }

  fmt.Println(persona1)
  fmt.Println(persona1.nom)
  fmt.Printf("%T", persona1)

  enginyer1 := enginyer{
    persona:persona{
      nom: "Isabel",
      edat: 36,
    },
    realitzaPlanols: true,
  }

  fmt.Println(enginyer1)

}
