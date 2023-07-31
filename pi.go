package main

import "fmt"

const NUM = 4

func powMinusOne(k int) int {
	if k%2 == 0 {
		return 1
	}
	return -1
}

func calculateNumbers(number int) float64 {
	s := 0.0
	for k := 1; k <= number; k++ {
		s += float64(powMinusOne(k+1)) / float64(2*k-1)
	}
	return float64(NUM) * s
}

func main() {
	fmt.Println(calculateNumbers(100000000))
}
