package main

import "fmt"

type Person struct {
  Name string
}

type Android struct {
  Person
  Model string
}

func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

func main(){
  p := Person{"jaime"}
  p.Talk()


  a := new(Android)
  a.Person.Talk()

  b := new(Android)
  b.Talk()
}
