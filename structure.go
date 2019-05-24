package main
import "fmt"

type Dog struct {
  Age int
  Name string
}

func (d *Dog) Guau() {
  fmt.Println("Guau Guau!", d.Name)
}

func main(){
  c := Dog{2,"Boby"}
  fmt.Println(c)
  c.Guau()

}
