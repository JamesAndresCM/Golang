package main

import (
    "fmt"
    "math"
)

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}

type Square struct {
    Width  float64
    Height float64
}

func (s Square) Area() float64 {
    return s.Width * s.Height
}

type Figure interface {
    Area() float64
}

func Print(f Figure, name string) {
  fmt.Printf("%s, Area result for %f is %f\n",name, f, f.Area())
}
func main() {
    c := Circle{Radius: 10}
    s := Square{Height: 10, Width: 5}
    Print(c, "Circle")
    Print(s, "Square")
}
