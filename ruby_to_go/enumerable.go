package main

import "fmt"

func main(){
  numbers := []int{1,2,3,4}
  fmt.Println(collect(numbers))
  animals := []string{"perro", "gato", "leon"}
  fmt.Println(any(animals, 4))
  fmt.Println(all(animals, 2))
  fmt.Println(cycle(animals, 2))
  fmt.Println(find(numbers, 2))
  fmt.Println(drop(numbers, 1))
}

func collect(data []int) []int{
  result := []int{}
  for i := 0; i < len(data); i++{
    result = append(result, data[i] * data[i])
  }
  return result
}

func any(data []string, length int) bool{
  result := []bool{}
  for _, val := range data {
    if len(val) > length{
      result = append(result, true)
    }
  }
  
  var cond bool
  for _, v := range result{
    if v == true {
      cond = true
      break
    }
  }
  return cond
}

func all(data []string, length int) bool{
  result := []bool{}
  for _, val := range data{
    if len(val) > length{
      result = append(result, true)
    } else {
      result = append(result, false)
    }
  }
  quantity := []bool{}
  for _,val := range result{
    if val == true {
      quantity = append(quantity, true)
    }
  }
  
  if len(data) == len(quantity){
    return true
  } else {
    return false
  }
}

func cycle(data []string, length int) []string{
  result := []string{}
  for i := 0; i < length; i++{
    for _, val := range data{
      result = append(result,val)
    }
  }
  return result
}

func find(data []int, num int) []int{
  result := []int{}
  for _, val := range data{
    if val > num {
      result = append(result, val)
    }
  }
  return result
}

func drop(data []int, num int) []int{
  drop := data[num:]
  return drop
}
