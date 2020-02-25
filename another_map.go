package main

import "fmt"

func main(){
  countries()
}

func countries(){
  teams := map[string]int{
    "Universidad de Chile": 30,
    "River Plate": 40,
    "Colo Colo": 42,
    "Barcelona": 79,
    "Real Madrid": 13}

  for index, value := range teams {
    fmt.Printf("Key: %s, Value: %d\n",index, value)
  }
}
