package main

import (
	"fmt"
)

func main(){
  //Un DataMap és el que seria un Object a JS
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


}
