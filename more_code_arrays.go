package main

import "fmt"

func main(){
  var data = [12]string{
    "hola","adios","auto","hola","hola","gato","perro","adios","go","auto","auto","perro"}

  hash := make(map[string]int)
  for i := 0; i < len(data); i++ {
    hash[data[i]] +=1 
  }
  fmt.Println(hash)
  men_may()
  fmt.Println("Result :",fibonacci(10))
}

func men_may(){
  var data = [10]int{10,1,23,23,102,34,56,12,344,100}

  min := data[0]
  max := data[0]

  for _, value := range data {
    if min < value {
      min = value
    }
    if max > value {
      max = value
    }
  }
  fmt.Println("MAX :",min)
  fmt.Println("MIN :",max)
}

func fibonacci(n uint) uint {
  if n <= 1 {
    return n
  }
  return fibonacci(n - 1) + fibonacci(n - 2)
}
