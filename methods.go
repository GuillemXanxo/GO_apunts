//Given some inputs they will change the state of those inputs (in contrast with funcs)

package main

import (
	"fmt"
)

type Person struct {
  FirstName string
  LastName string
  Age int
  DNI string
  Status bool
}
//VALUE RECEIVER METHOD
func (p Person) Identification() string {
  return fmt.Sprintf("%s %s is age %d", p.FirstName, p.LastName, p.Age)
}
/*(p Person) is the receiver and Identification is the name of the method
Hem creat un metodo a struct Person, mirar al main com sutilitza
*/

//POINTER RECEIVER METHOD --> this actually make changes to the struct
func (p *Person) SetDNI(dni string) {
  p.DNI = dni 
}

func (p Person) GetStatus() bool { //No serveix de res fer getters ni setter, es recomanable fer-ho al valor directament
  return p.Status
} 

func main() {
  character := Person {
    FirstName: "Guillem",
    LastName: "Xnx",
    Age: 32,
    Status: true,
  }

  output := character.Identification()
  fmt.Println(output)

  character.SetDNI("12345678C") // en realita GO fa (&character).DNI --> esta convrtint a pointer
  fmt.Println(character)

  if character.GetStatus() {
    fmt.Println("Character is online")
  } else {
    fmt.Println("Character is offline")
  }

  if character.Status {
    fmt.Println("Character is connected")
  } else {
    fmt.Println("Character is disconnected")
  }
}
