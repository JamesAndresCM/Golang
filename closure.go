package main

import "fmt"

func main() {
  num := num_value()
  five := Mult(num)
  for i := 1; i < 11; i++ {
    con := fmt.Sprintf("%d * %d = %d",i ,num ,five())
    fmt.Println(con)
  } 
}

func num_value() int {
  var num int 
  fmt.Println("Entry num :")
  _, err := fmt.Scanf("%d", &num) 
  if err != nil{
    return 0
  }
  return num
}

func Mult(val int) func() int {
  num := val
  seq := 0
  return func() int {
    seq +=1
    return num * seq
  }
}

