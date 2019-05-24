package main

import "fmt"

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func min_max(numbers []int) (int,int) {
  min := numbers[0]
  max := numbers[0]

  for _, value := range numbers {
    if value < min {
      min = value
    } 
    if value > max {
      max = value
    }
  }
  return min, max
}

func main(){
  var numbers int
  elements := []int{}
  var val int

  fmt.Printf("Entry quantity numbers: \n")
  fmt.Scanf("%d\n", &numbers)
  
  for i := 0; i < numbers; i++ {
    fmt.Println("Entry Number:", i)
    fmt.Scanf("%d\n", &val)
    elements = append(elements, val)
  }
  fmt.Println(elements)

  reverse(elements)
  fmt.Println("Reverse Array",elements)
  fmt.Println(min_max(elements))
}

