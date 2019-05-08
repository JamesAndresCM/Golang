package main

import  (
  "fmt"
  "bufio"
  "os"
)

func main(){
  age := 0
  name := ""
  fmt.Println("Entry your age: ")
  fmt.Scanf("%d\n",&age)
  fmt.Printf("Your age is %d\n", age)
  fmt.Println("Entry your name: ")
  fmt.Scanf("%s\n",&name)
  fmt.Printf("Your name is %s\n", name)

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Entry hobby : ")
  text,err := reader.ReadString('\n')

  if err != nil {
    fmt.Println(err)
  }else{
    fmt.Println("Your hobby is: ", text)
  }
}
