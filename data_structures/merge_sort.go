package main

import "fmt"

func main() {
	arr := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 1, 100, 5, 23}
	fmt.Println(merge_sort(arr))
}

func merge_sort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	arr_size := len(arr)
	mid := arr_size / 2
	left := arr[0:mid]
	right := arr[mid:arr_size]

	left = merge_sort(left)
	right = merge_sort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	last_left := left[len(left)-1]
	first_right := right[0]

	if last_left <= first_right {
		return append(left, right...)
	}

	result := make([]int, 0)
	var min int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			min = left[0]
			left = left[1:]
		} else {
			min = right[0]
			right = right[1:]
		}
		result = append(result, min)
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
