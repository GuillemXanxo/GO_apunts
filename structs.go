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
  persona1 := persona{ //primera manera de declarar un struct literal
    nom: "Guillem",
    edat: 32,
  }

  persona2 := persona{ //segona manera de declarar un struct literal
    "Maria",
    26,
  }

  fmt.Println(persona1)
  fmt.Println(persona1.nom)
  fmt.Printf("%T", persona1) // imprimeix el tipus de dada que es persona1, en aquest cas es persona declarat a main 
  fmt.Println(persona2)

  enginyer1 := enginyer{
    persona:persona{
      nom: "Isabel",
      edat: 36,
    },
    realitzaPlanols: true,
  }

  fmt.Println(enginyer1)
  fmt.Println(enginyer1.realitzaPlanols) //podem accedir als camps de persona des de enginyer
  fmt.Println(enginyer1.nom)

  //ANONYMOUS STRUCTS (utils per mashall i unmarshall data d'un JSON)
  //Primera manera
  var person struct {
    name string
    age int
    pet string
  }

  person.name = "Bob"
  person.age = 50
  person.pet = "dog"

  //Segona manera
  pet := struct{
    name string
    kind string
  }{
    name: "Becker",
    kind: "dog",
  }

  fmt.Println(pet)
}
