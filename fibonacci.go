package main

import "fmt"

func main(){
  fmt.Printf("Result : %d\n", fibonacci(10))
}

func fibonacci(num int) int{
  if num <=1 {
    return num
  }
  return (fibonacci(num - 1) + fibonacci(num - 2))
}
