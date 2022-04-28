// https://codeo.app/problemas/0xcf-mirando-al-horizonte
package main

import "fmt"

func main() {
	roofs := []int{4, 2, 1, 6, 3, 2, 5, 3, 1, 4}
	fmt.Println(solve(roofs))
	roofs = []int{1, 1, 2}
	fmt.Println(solve(roofs))
	roofs = []int{3, 3, 2, 2, 1, 1}
	fmt.Println(roofs)
}

func solve(roofs []int) []int {
	result := []int{}
	var n int
	for i := 0; i < len(roofs); i++ {
		n = -1
		for j := i + 1; j < len(roofs); j++ {
			if roofs[i] < roofs[j] {
				n = roofs[j]
				break
			}
		}
		result = append(result, n)
	}
	return result
}
