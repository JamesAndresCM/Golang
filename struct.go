package main

import "fmt"

type User struct{
  age int
  name, last_name string
}

func (this *User) to_s() string{
  return this.name + " " + this.last_name
}

func main(){
  user_ := new(User)

  user_.name = "james"
  user_.last_name = "canales"

  fmt.Println(user_.to_s())
  
}
