package main

import "fmt"

func main(){
  numbers := []int{1,2,3,4,50,65,101,400}
  fmt.Println(numbers)
  var num int
  fmt.Println("Entry number to search:")
  fmt.Scanf("%d", &num)

  if num != 0 {
    result := bynarySearch(numbers, num)
    if result < 0 {
      fmt.Printf("Number %d not found\n", num)
    } else {
      fmt.Printf("Position of number %d\n",result)
    }
  }
}

func bynarySearch(arr []int, element int) int{
  result := -1
  i := 0
  j := len(arr) - 1

  for i <= j {
    middle := (i + j) / 2
    if arr[middle] == element {
      result = middle
      break
    } else if arr[middle] < element {
      i = middle + 1
    } else if arr[middle] > element {
      j = middle - 1
    }
  }
  
  return result
}
