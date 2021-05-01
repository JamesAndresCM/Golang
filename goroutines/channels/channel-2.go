package main

import "fmt"

func main(){
  ch := make(chan string, 6)
  go func(msg string){
    for i := 0; i < 6; i++{
      fmt.Printf("Sending: %s - %d\n", msg, i)
      ch <- msg
    }
    close(ch)
  }("hello world")

  for v := range ch {
    fmt.Println("Received: ",v)
  }
}
