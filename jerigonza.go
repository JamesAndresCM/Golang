package main

import (
	"fmt"
	"strings"
)

var vocals = []string{"a", "e", "i", "o", "u"}

func main() {
	word := "hola como estas hoy en dia"
	fmt.Println(convert(word))
	word2 := "hola"
	fmt.Println(convert(word2))
	word3 := "beyond two souls"
	fmt.Println(convert(word3))
}

func convert(word string) string {
	split := strings.Split(word, "")
	data := []string{}

	for _, val := range split {
		for _, res := range vocals {
			if res == val {
				k := val + "p"
				data = append(data, k)
			}
		}
		data = append(data, val)
	}
	return strings.Join(data, "")
}
