package main

import "fmt"

func Increase(num *int) int{
  *num *= 2
  return *num
}

func Inc(num *int){
  *num++
}

func Sum(num int) int{
  return num + 2
}
func main(){
  val := 6
  Increase(&val)
  fmt.Println(val)

  val2 := 4
  res2 := Sum(val2)
  fmt.Println(res2)

  val3 := 2
  Inc(&val3)
  fmt.Println(val3)
}
