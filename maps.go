package main

import "fmt"


func main(){
  x := make(map[string]int)
  x["a"] = 10
  x["b"] = 20
  x["c"] = 30
  x["d"] = 40
  delete(x,"a")
  fmt.Println(x)
  create_dynamic()
  another_func()
}

func create_dynamic(){
  hash := make(map[string]int)
  letters := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","t","u","v","w","x","y","z"}

  for index, value := range letters {
    hash[value] = index
  }
  fmt.Println(hash)
}

func another_func(){
  hash := make(map[string]int)
  hash["a"] = 1
  hash["b"] = 2
  hash["c"] = 3
  hash["d"] = 4

  if name, ok := hash["a"]; ok {
    fmt.Println(name, ok)
  }
}
