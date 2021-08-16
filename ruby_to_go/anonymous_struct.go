package main
import "fmt"

func main(){
  dog := struct {
    Name string
    Country string
  }{"Cholo", "Chile"}

  fmt.Printf("Dog name: %s, Country: %s\n",dog.Name, dog.Country)


  cat := struct {
    int
    string
  }{2, "Cuvi"}
  fmt.Printf("Cat year: %d, Name: %s\n",cat.int, cat.string)
}
