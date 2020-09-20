package main

import "fmt"

func main() {
	original_arr := []int{412, 12, 123, 1, 2324, 3, 155, 23412, 2}
	fmt.Println(original_arr)
	order_arr := bubble_sort(original_arr)
	fmt.Println(order_arr)
}

func bubble_sort(arr []int) []int {
	orders := 0
	change := true
	for change {
		change = false
		for i := 1; i < (len(arr) - orders); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				change = true
			}
		}
		orders++
	}
	return arr
}
