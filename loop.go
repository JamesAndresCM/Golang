package main

import "fmt"

func main(){
  sum := 0

  for i := 0; i < 10; i++ {
    sum+=i
  }
  fmt.Println(sum)

  while()
  do_while()
}

func while(){
  i := 0
  for i < 10 {
    i++
    fmt.Println("while : ",i)
  }
}

func do_while(){
  i := 0
  for i < 10{
    i++
    if i < 3 {
      fmt.Println("do_while",i)
    }
  }
}
