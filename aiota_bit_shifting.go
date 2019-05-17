package main

import "fmt"


const (
  _ = iota
  kb = 1 << (iota * 10)
  gb = 1 << (iota * 10)
  tb = 1 << (iota * 10)
  
)

func main(){
  var a int = 2
  fmt.Println(a << 3)
  fmt.Println("Kilobyte :",kb, " bytes")
  fmt.Println("Gigabyte :",gb, " bytes")
  fmt.Println("Terabyte :",tb, " bytes")
}
