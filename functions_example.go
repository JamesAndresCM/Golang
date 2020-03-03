package main

import "fmt"

const temp = 5
type Months int

const (
  January = iota + 1
  February
  March
  April
  May
  June
  July
  August
  September
  October
  November
  December

)
func isZERO(n int) bool { return n == 0  }

func getDAY(n int) {
  days := map[int]string{
    1: "Lunes",
    2: "Martes",
    3: "Miercoles",
    4: "Jueves",
    5: "Viernes",
    6: "Sabado",
    7: "Domingo",
  }
  fmt.Println(days[n])
}

func split(word string) {
  var data []rune
  for _, a := range word {
    data = append(data, a)
  }
  fmt.Printf("%q\n",data)
}

func join(x []string) {
  for _, a := range x {
    fmt.Printf("%s",a) 
 }
 fmt.Println()
}

func main() {
  a := temp
  result := a * 2
  var number int = 10
  if isZERO(number) {
    fmt.Printf("Number %d is Zero %t\n",number)
  }
  
  fmt.Printf("Number %d is not Zero\n",number)
  
  fmt.Println(result)
  fmt.Println(January)
  getDAY(1)

  months := [...]string{
                        January: "Enero",
                        February: "Febrero",
                        March: "Marzo",
                        April: "Abril",
                        May: "Mayo",
                        June: "Junio",
                        July: "Julio",
                        August: "Agosto",
                        September: "Septiembre",
                        October: "Octubre",
                        November: "Noviembre",
                        December: "Diciembre",
                      }
  fmt.Printf("Nº: %d, Name: %s\n",December, months[December])
  split("nameless")
  z := []string{"h","e","l","l","o"}
  join(z)
}
