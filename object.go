package main

import "fmt"

type Person struct {
  Name string
  Age int
  City string
  Hobbie string
}

func (p *Person) CalcAge(number int) string{
  if number > 20 {
    return "Mayor de Edad"
  }
  return "Menor de Edad"
}

func (p *Person) Deportes(names ...string) map[string]int {
  data := make(map[string]int)
  for i, value := range names {
    data[value] = i
  }
  return data
}

func (p *Person) ToString() {
  fmt.Printf("Hello I'm: %v, Age: %v, City: %v, Hobbie: %v\n", p.Name, p.Age, p.City, p.Hobbie)
}

func (p *Person) getRegion(name string) {
  data := make(map[string]int)

  data["Santiago"] = 0
  data["Talca"] = 1
  data["Temuco"] = 2
  data["Arica"] = 3

  for x,y := range data {
    if name == x {
      fmt.Printf("Tu Ciudad es : %v, Numero : %v\n", x, y)
    }
  }
}

func main(){
  p := Person{"Jhon Doe", 30, "Santiago", "Read"}
  p.ToString()
  fmt.Println(p.CalcAge(p.Age))
  deportes := p.Deportes("tenis","natacion", "futbol")
  fmt.Println(deportes)
  p.getRegion(p.City)
}
