//Sentinel errors should be rare, it is better to use already existing ones

package main

import (
  "fmt"
  "errors" //Importem package errors
)
 //Creem els nostres errors propis que som els que retornarem per fer mes facil la comprensio de lerror
var (
  ErrHourlyRate = errors.New("invalid hourly rate")
  ErrHoursWorked = errors.New("invalid worked hours per week")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
  if hourlyRate < 10 || hourlyRate > 75 {
    return 0, ErrHourlyRate
  }
  if hoursWorked < 0 || hoursWorked > 80 {
    return 0, ErrHoursWorked 
  }
  if hoursWorked > 40 {
    hoursOver := hoursWorked - 40
    overTime := hoursOver * 2
    regularPay := hoursWorked * hourlyRate
    return regularPay + overTime, nil
  }
  return hoursWorked * hourlyRate, nil
}


func main() {
 pay, err := payDay(81, 50) //error hoursworked
 if err != nil{
   fmt.Println(err)
 }
//Si aqui impreixo pay em donara 0

 pay, err = payDay(80, 5) //error hourlyrate
 if err != nil{
   fmt.Println(err)
 }
 
 
 pay, err = payDay(80, 50)
 if err != nil{
   fmt.Println(err)
 } 
 fmt.Println(pay)
}
