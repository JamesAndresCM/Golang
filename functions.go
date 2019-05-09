package main

import "fmt"

func main(){
  data := []int{1,234,12,34,444,123}
  fmt.Println(calc(data))
  fmt.Println(mult_args(123,4,12,12233,23,45,6,75,677,45,66))

  closure := func(x,y int) int{
    return x + y  
  }
  fmt.Println(closure(10,20))
  fmt.Println(factorial(6))
}

func calc(x []int) int {
  total := 0
  for _, value := range x {
    total += value
  }
  return total / len(x)
}

func mult_args(args ...int) []int {
  data := []int{}
  for _, value := range args {
    if value % 2 == 0 {
      data = append(data, value)
    }
  }
  return data
}

func factorial(x uint) uint { 
  if x == 0 {
    return 1 
  }
  return x * factorial(x-1) 
}
