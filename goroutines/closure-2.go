package main

import (
  "sync"
  "fmt"
)
func main(){
  var wg sync.WaitGroup
  for i := 1; i <= 3; i++{
    wg.Add(1)
    go func(i int){
      defer wg.Done()
      fmt.Println(i)
    }(i)
  }
  wg.Wait()
}
