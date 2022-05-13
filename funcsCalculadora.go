package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int {return i + j}
func sub(i int, j int) int {return i - j}
func mul(i int, j int) int {return i * j}
func div(i int, j int) int {return i / j}

var operatorsMap = map[string]func(int, int) int {
  "+": add,
  "-": sub,
  "*": mul,
  "/": div,
}

func main(){
  expressions := [][]string{
    []string{"2","+","3"},
    []string{"2","-","3"},
    []string{"2","*","3"},
    []string{"2","/","3"},
    []string{"2","%","3"},
    []string{"two","+","three"},
    []string{"5"},
  }

  //FIXA'T EN LA GESTIO D'ERRORS, ES LA CLAU! 
  for _, expression := range expressions {
    if len(expression) !=3 {
      fmt.Println("invalid expression: ", expression)
      continue
    }
    p1, err := strconv.Atoi(expression[0]) // convertim el string en [0] a int
    if err != nil {   //si hi ha error imprimeix error i continua el loop
      fmt.Println(err)
      continue
    }
    operator := expression[1]
    operatorFunction, ok := operatorsMap[operator] //Busquem operador a [1] en el map de funcions
    if !ok { //si no rebem el ok == true aleshores salta i continua el loop
      fmt.Println("unsupported operator: ", operator)
      continue
    }
    p2, err := strconv.Atoi(expression[2]) // convertim el string en [2] a int
    if err != nil {
      fmt.Println(err)
      continue
    }
    result := operatorFunction(p1, p2)
    fmt.Println(result)
  }
}
