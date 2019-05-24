package main

import "fmt"

func half(number int)(int, bool){
  if number % 2 == 0 {
    return 1, true
  }else {
    return 0, false
  }
}

func find_fibonacci(number int) int{
  if number == 0 {
    return 0
  } else if number == 1 {
    return 1
  }else{
    return find_fibonacci(number - 1) + find_fibonacci(number - 2)
  }

}


func max_number(number ...int) int{
  max := number[0]
  for _, value := range number {
    if max > value {
      max = value
    } 
  }
  return max
}

func main(){
  fmt.Println(half(1))
  fmt.Println(half(2))
  fmt.Println(max_number(11,34,3))
  fmt.Println(find_fibonacci(14))
}
