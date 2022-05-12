package main

import (
	"fmt"
)

func main(){
  
  m := map[string]int{
    "Josep": 6,
    "Maria": 8,
  }
  fmt.Println(m)

  //DataMaps anidats:
  mapEixample := map[string]int{
    "EixampleDreta": 85000,
    "EixampleEsquerra": 92000,
  }

  mapCiutatVella := map[string]int{
    "Raval": 45000,
    "Born": 33000,
  }

  mapBarcelona := map[string]map[string]int{
    "Eixample": mapEixample,
    "CiutatVella": mapCiutatVella, 
  }

  fmt.Println(mapBarcelona["Eixample"])

  //Map de slices
  equiposLFP := map[string][]string{
    "Barcelona": []string{"Ter Stegen", "Pique", "Messi"},
    "Madrid": []string{"Courtois", "Hierro", "Benzema"},
    "Espanyol": []string{"RDT", "Perosa", "Melendo"},
  }

  fmt.Println(equiposLFP)
  fmt.Println(equiposLFP["Barcelona"])
  fmt.Println(equiposLFP["Barcelona"][1])

  equiposLFP["Barcelona"][1] = "Zubizarreta"
  fmt.Println(equiposLFP["Barcelona"][1])

  // How to know if a key exists in a map
  m1 := map[string]int{
    "hello": 5,
    "world": 0,
  }
  v1, ok1 := m1["hello"] //v sera el nom de la key y ok sera true si existeix
  fmt.Println(v1, ok1)

  v2, ok2 := m1["goodbye"]
  fmt.Println(v2, ok2)

  //Delete a key from a map (using m1 from above)
  delete(m1, "world") // delete receives de map we target and the key to be deleted
  fmt.Println(m1)


}
