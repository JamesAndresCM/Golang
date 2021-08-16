package main

import "fmt"

type ProduceBasket interface {
	AddItem(entry Produce) Basket
	RemoveItem(entry Produce) Basket
	Items()
}

type Basket []Produce
type Produce string

func (p *Basket) AddItem(entry Produce) Basket {
	*p = append(*p, entry)
	return *p
}

func (p *Basket) RemoveItem(entry Produce) Basket {
	result := *p
	for i, v := range result {
		if v == entry {
			result = append(result[:i], result[i+1:]...)
			break
		}
	}
	*p = result
	return result
}

func (p Basket) Items() {
	for _, v := range p {
		fmt.Println(v)
	}
}

func main() {
  fruits := new(Basket)
	fruits.AddItem("Apple")
	fruits.AddItem("Mango")
  fruits.RemoveItem("Mango")
	
  veggies := new(Basket)
  veggies.AddItem("Broccolli")

  var items []ProduceBasket
  items = append(items, fruits)
  items = append(items, veggies)

  for _, v := range items{
    v.Items()
  }
}
