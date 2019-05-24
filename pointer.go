package main

import "fmt"

func pointer(x *int){
  *x = 0
}

func main(){
  x := 5
  pointer(&x)
  fmt.Println(x)
}
