package main

import (
  "fmt"
  "time"
  user "./user"
)
// inheritance
type person struct {
  user.User
}

func main() {
  userTest := new(user.User)
  userTest.Initialize(1, "Jhon", "Doe", time.Now(), true)
  fmt.Println("First User:",userTest)
  fmt.Println("CompleteName:", userTest.ToS())
  fmt.Println("\n")

  userD := user.User{1, "Test", "another", time.Now(), false}
  fmt.Println("Second User:",userD)
  fmt.Println("CompleteName:",userD.ToS())
  fmt.Println("\n")

  userI := new(person)
  userI.Initialize(1, "Inheritance", "Example", time.Now(), true)
  fmt.Println("Third User:",userI.User)
  fmt.Println("CompleteName:",userI.ToS())
}
