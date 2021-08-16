package main

import "fmt"

type basket []product

type product struct{
  name string
  flavour string
  kind string
}

func (p *basket) addItem(entry product){
  *p = append(*p, entry)
  fmt.Println(entry.name + " Created")
}

func (p basket) changeItem(name string, entry product){
  for key, val := range p {
    if val.name == name{
      p[key] = entry
    }
  }
  fmt.Println("Product changed" + name)
}


func (p basket) listItems(){
  fmt.Println("\nItems\n***")
  for _, v:= range p {
    fmt.Println("Product name: " +v.name)
    fmt.Println("Product flavour: " +v.flavour)
    fmt.Println("Product kind: " +v.kind)
    fmt.Println()
  }
}

func main(){
  basket := basket{}
  basket.addItem(
  product{
    name: "apple",
    flavour: `It's a little sour and bitter, but mostly sweet,notatallsalty,veryjuicyingeneral.`,
    kind: "ira",
    },
  )

 basket.addItem(
  product{
    name: "carrot",
    flavour: `It's a little sour and bitter, but mostly sweet,notatallsalty,veryjuicyingeneral.`,
    kind: "ira2",
    },
  )

  basket.changeItem("carrot",
    product{
      name: "carrot",
      flavour: "prueba",
      kind: "hola",
    },
  )

  basket.listItems()

}
