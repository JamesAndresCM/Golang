package main

import (
  "fmt"
  "reflect"
)

func main(){
  letters := []string{"a","b","c"}
  letters = append(letters, "d")
  fmt.Println(letters)
  

  items("hello","world")
  fruits := []string{"apple","banana", "orange", "mango"}
  items(fruits...)

  numbers := []int{1,2,3,4,5,12,34}
  fmt.Println(removeItem(numbers, 12))
  fmt.Println(removeItem(fruits, "mango"))
  arrInterface := []interface{}{1,"james", 2123}
  fmt.Println(removeItem(arrInterface, "james"))
}

func items(args ...string){
  for _, item := range args{
    fmt.Println("output:",item)
  }
}


func removeItem(arr interface{}, el interface{}) []interface{}{
  var result []interface{}
  data := reflect.ValueOf(arr).Index(0).Kind().String()
  switch data {
  case "string":
    for _, index := range arr.([]string){
      if index != el {
        result = append(result, index)
      }
    }
  case "int":
    for _, index := range arr.([]int){
      if index != el {
        result = append(result, index)
      }
    }
  case "interface":
    for _, index := range arr.([]interface{}){
      if index != el {
        result = append(result, index)
      }
    }
  default:
    fmt.Println("unknown", data)
  }
  return result
}
