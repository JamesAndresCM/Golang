package main
import "fmt"

func main(){
  fruit := map[string]string{
      "name": "apple",
      "flavour": "test",
  }
  basket(
   fruit,"fruit",
  )
}

func basket(args ...interface{}){
  item := args[0].(map[string]string)
  kind := args[1]

  fmt.Printf("Name: %s, Flavour: %s, Kind: %s\n", item["name"], item["flavour"], kind)
}
