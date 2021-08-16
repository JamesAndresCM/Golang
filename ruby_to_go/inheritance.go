package main

import "fmt"
type Animal struct {
  Kind string
  Diet string
}

type Owner struct {
  Name string
  Country string
}
type Dog struct {
  *Animal
  *Owner
  Name string
  Breed string
}

type Cat struct {
  *Animal
  *Owner
  Name string
  Breed string
}

func main(){
  dog := Dog{
    Name: "Cholo",
    Breed: "kilter",
  }
  dog.Animal = &Animal{
    Kind: "Dog",
    Diet: "Omnivorous",
  }
  dog.Owner = &Owner{
    Name: "James",
    Country: "Chile",
  }
  
  fmt.Printf("Name %s, Breed: %s\n",dog.Name, dog.Breed)
  fmt.Printf("Animal Kind: %s, Diet: %s\n",dog.Animal.Kind, dog.Animal.Diet)
  fmt.Printf("%s live in %s\n", dog.Owner.Name, dog.Owner.Country)
}

