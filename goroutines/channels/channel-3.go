package main
import "fmt"

func sendMsg(msg string, ch1 chan <- string) {
  ms := "pong!!!!"
  ch1 <- msg + "*****" + ms
}

func receiveMsg(ch1 <- chan string, ch2 chan <- string) {
  m := <-ch1
  ch2 <- m
}
func main(){
  ch1 := make(chan string)
  ch2 := make(chan string)

  go sendMsg("ping", ch1)

  go receiveMsg(ch1, ch2)

  result := <-ch2
  fmt.Println(result)
}
