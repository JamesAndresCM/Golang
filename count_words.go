package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	t := make(map[string]int)
	words := strings.Fields(s)
	for i := range words {
		t[words[i]]++
	}
	return t
}

func main() {
	fmt.Println(WordCount("Im learn learn to go go"))
}
