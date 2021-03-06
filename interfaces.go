package main

import (
	"fmt"
	"math"
)

type circle struct {
  radius float64
}

type square struct {
  width float64
  height float64

}
// both structures have a method to calculate the area:

func (s square) area() float64 {
  return s.height * s.width
}

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}

/* Per tant, circle i square tenen comportaments similars.
Una interface ens permet relacionar els elements amb un comportament similar */

type shape interface {
  area() float64
} // --> qualsevol struct o type que tingui el metode area() es de interface shape
//Es com un tipus superior del struct. Si square i circle poden tenir area es que els dos son formes geometriques
//POLYMORPHISM: a shape can ppear as a circle or square. Si una funcio agafa com a parametre un shape vol dir
//que agafa circle i square --> hem fet una funcio que escapa als tipus concrets a traves de polimorfisme

func main() {
  circle1 := circle{4.5}
  square1 := square{3, 5.5}
  /* Vull calcular area de totes les formes geometriques:
  shapes := []type{circircle1, square1} --> pero de quin tipus de [] es? aqui entra en joc interface 
  Definim la interface shape, que ens donara el tipus per slice */

  shapes := []shape{&circle1, &square1}
  for _, shape := range shapes {
    fmt.Println(shape.area()) //no podem demanar res mes que el que tenen en comu, area()
  }

/* es bona practica quan fem el slice passar el & ja que si el method crida al punter funcionara igual
en canvi si no el posem, no funcionara 
UTILITAT
una interface pot ser acceptada com a parametre duna funcio. exemple la func totalArea */

  
  fmt.Println(totalArea(&circle1, &square1))

}

func totalArea(shapes ...shape) float64 { //Ens dona totes les arees del params que li passem
    var area float64
    for _, s := range shapes {
      area += s.area()
    }
    return area 
  }
