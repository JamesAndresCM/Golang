package main

import "fmt"

type Gender interface{
  sexHuman() string
}

type Alive interface{
  isAlive() bool  
}

type Human struct {
  age int
  name string
  sex string
  live bool
}

type Men struct {
  Human
}

type Women struct {
  Human
}

func (h *Human) sexHuman() string {
  if h.sex == "men" {
    return "hombre"
  } else {
    return "mujer"
  }
}

func (h *Human) isAlive()  bool {
  return h.live
}

func GetGender(g Gender){
  fmt.Printf("Gender: %s\n", g.sexHuman())
}

func GetLiveStatus(l Alive){
  if l.isAlive() {
    fmt.Printf("Status: LIVE\n") 
  }else{
    fmt.Printf("Status: DEAD\n")
  }
}


func main(){
  male := new(Human)
  male.name = "jhon"
  male.sex = "men"
  fmt.Printf("Name: %s, ", male.name)
  GetLiveStatus(male)
  GetGender(male)

  female := new(Human)
  female.age = 20
  female.name = "mary"
  female.live = true
  fmt.Printf("Name: %s, ", female.name)
  GetGender(female)
  GetLiveStatus(female)
}
