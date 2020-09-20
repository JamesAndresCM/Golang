package main

import "fmt"

func main() {
	arr := []int{31, 41, 59, 26, 53, 58, 97, 93, 23, 84, 1, 11, 200}
	quickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(array *[]int, left, right int) {
	if left < right {
		arr := *array
		limit, origin := right, left
		pivot := arr[right]
		right--

		for left <= right {
			for left < len(arr) && arr[left] < pivot {
				left++
			}

			for right >= 0 && arr[right] > pivot {
				right--
			}

			if left <= right {
				if arr[left] > arr[right] {
					arr[left], arr[right] = arr[right], arr[left]
				}
				left++
				right--
			}

		}
		swap(array, left, limit)
		quickSort(array, origin, right)
		quickSort(array, left, limit)
	}

}

func swap(array *[]int, left, right int) {
	arr := *array
	tmp := arr[left]
	arr[left] = arr[right]
	arr[right] = tmp
}
