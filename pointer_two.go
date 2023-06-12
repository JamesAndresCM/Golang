package main

import (
  "fmt"
  "strings"
  )

const days = 7


type Child struct {
  Name string
  Age int
}

func (child Child) toString() {
  fmt.Println("nombre y edad ", child.Name, child.Age)
}

func (child Child) mayorEdad(edad int) {
  if edad > 18 {
    fmt.Println("es mayor de edad!")
   } else {  
    fmt.Println("no es mayor de edad!")
   }
}

func cambiarEdad(child *Child, newAge int) Child{
  child.Age = newAge
  return *child
}

func cambiarNombre(child Child, newName string) Child {
  child.Name = newName
  fmt.Println(child.Name)
  fmt.Println(child)
  return child
}

func main() {
  var age int
  age = 10
  word := "james"
  data := []int{1,2,3,4,5}
  fmt.Printf("your age is %d\n", age);
  fmt.Printf("days is %d\n", days);
  sumArray := iterateElements(data)
  fmt.Println(sumArray)
  fmt.Println(strings.Contains(word, "am"))


  for count := 1; count < 10; count++ {
      fmt.Println(count)
  }

  hija := Child{Name: "Mia Elena", Age: 3}
  hija.toString()
  hija.mayorEdad(23)

  cambiarEdad(&hija, 80)
  fmt.Println(hija.Name, hija.Age)
  
  cambiarNombre(hija, "Ira")
  fmt.Println(hija.Name, hija.Age)
}

func iterateElements(elements []int) int{ 
  var result int
  for _,i := range elements {
    result += i
  }
  return result

} 
