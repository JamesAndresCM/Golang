package main

import "fmt"

type Person struct {
  Name string
  Age int
}

type Album struct {
  Name string
}

type Disc struct {
  Person
  Title string
}

func (a *Album) ListSongs(){
  data := make(map[string]int)
  data["primer_tema"] = 1
  data["segundo_tema"] = 2
  data["tercer_tema"] = 3
  fmt.Printf("Album, Canciones %v%v\n", a, data)
}


func main(){
  song := Album{"Corazones Rojos"}
  song.ListSongs()
  
  person := Person{"Jhon Doe", 30}
  disc := Disc{person, "name one"}

  fmt.Printf("Autor : %v\n",disc)

}
