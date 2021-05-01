package main
import (
  "fmt"
  "time"
)

func main() {
  fun("hello")

  go fun("hello-2")

  go func(){
    fun("hello-3")
  }()

  vt := fun
  vt("hello-4")
  fmt.Println("waiting for goroutines")
  time.Sleep(100 * time.Millisecond)
}

func fun(s string){
  for i := 0; i < 3; i++{
    fmt.Println(s)
    time.Sleep(1 * time.Millisecond)
  }
}
