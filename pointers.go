package main

import "fmt"


func main(){

  var x,y *int
  pointer := 5

  x = &pointer
  y = &pointer

  *x = 6

  fmt.Println(*x)
  fmt.Println(*y)
}
