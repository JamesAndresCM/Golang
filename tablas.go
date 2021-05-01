package main

import "fmt"

func main(){
  pares := []int{}
  impares := []int{}

  var n int
  fmt.Println("Ingresa un numero para multiplicar")
  fmt.Scanf("%d",&n)
  if n == 0 {
    fmt.Println("el numero ingresado debe ser mayor a 0")
    return
  }
  for i := 1; i < 11; i++{
    var result int
    result = i * n
    fmt.Println(n," x ", i, "es igual a: ", result) 
    if result % 2 == 0{
      pares = append(pares, result)
    }else {
      impares = append(impares, result)
    }
  }
  fmt.Println("numeros impares: ",impares)
  fmt.Println("numeros pares: ", pares)
}
