package main

import "fmt"

func main() {
	total := []int{}
	next := 1
	current := 0
	tmp := 0
	for i := 1; i < 20; i++ {
		tmp = current
		current = next
		next = next + tmp
		total = append(total, next)
	}
	fmt.Println(total)
}
