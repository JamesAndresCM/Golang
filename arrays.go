package main

import "fmt"

func main(){
  array := [5]int{1,2,3,4,5}

  for i := 0; i < len(array); i++{
    fmt.Println("Index:",i, "Value:",array[i])
  }
  another_loop()

}

func another_loop(){
  arr := [10]int{1,2,3,4,5,6,7,8,9,10}

  for i,value := range arr {
    fmt.Println("Pos: ",i, "Value: ",value)
  }
}
