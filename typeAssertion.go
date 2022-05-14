package main

import (
	"fmt"
)

func main() {
  var text interface{} = "I am a concrete type string"
  text2 := text.(string)//Controla si el tipus concret de text es string, si ho es passa el valor a text2
  //i si no es string?
  //text3 := text.(int) 
  //genera un panic q text no es int, i per tant no pot copiar el valor a text3 --> per aixo esta comentat
  //com ho fem?

  text3, ok := text.(int)
  if ok == false {
    fmt.Println(text3)
  }
  // la assertion torna tambe un boolean que si no compleix el tipus es false i a text3 assigna valor 0
  
}
