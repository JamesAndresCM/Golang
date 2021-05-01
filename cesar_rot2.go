package main

import (
  "fmt"
  "strings"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "ñ","o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func find(x string) int{
  result := 0
  for i, n := range letters {
        if x == n {
          result = i
        }
    }
  return result
}


func convert(val string) []int{
  data := []int{}
  value := strings.Split(strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(val, ",", ""), " ", "")), "")

  for i := 0; i < len(value); i++ {
    index := find(value[i])
    if value[i] == "z" {
      data = append(data, index + 1)
    }else if value[i] == "y" {
      data = append(data, index)
    }else {
      data = append(data, index + 2)
    }
  }
  return data
}

func main(){
  var value string
  fmt.Println("Entry text to convert rot2")
  fmt.Scanf("%s", &value)

  if len(value) == 0 {
    fmt.Println("Error, you must have a val")
    return
  }
  final := convert(value)

  con := []string{}
  for i,_ := range final {
    con = append(con, letters[final[i]])
  }
  fmt.Println("val:", value,",result:", strings.Join(con, ""))
}
