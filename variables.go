package main

import (
  "fmt"
  "strconv"
)

func main(){
  name := "james"
  age := 30
  eyes := "2"
  eyes_int,_ := strconv.Atoi(eyes)

  fmt.Println("name: " + name + " edad :" + strconv.Itoa(age) + " eyes : " , eyes_int + 2)

}
