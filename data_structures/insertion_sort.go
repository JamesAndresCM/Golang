/*
  toma el primer elemento de la izquierda y realiza la comparacion
  se asume que el primero esta ordenado
  [23,4,1,123,5]
  primer elemento = 4
  4 < 23
  [4,23,1,123,5]
  primer elemento = 1
  1 < 23 && 1 < 4
  [1,4,23,123,5]
  primer elemento = 5
  5 < 123 && 5 < 23 && 5 < 4 
  [1,4,5,23,123]
*/

package main
import "fmt"

func main() {
  original_arr := []int{412,12,123,1,2324,3,155,23412,2} 
  fmt.Println(original_arr)
  order_arr := insertion_sort(original_arr)
  fmt.Println(order_arr)
  order_arr_two := insertion_sort_two(original_arr)
  fmt.Println(order_arr_two)
}

func insertion_sort(arr []int) []int {
  for i := 1; i < len(arr); i++ {
    for j := 0; j < i; j++ {
      if arr[j] > arr[i] {
        arr[j], arr[i] = arr[i], arr[j]
      }
    }
  }
  return arr
}

func insertion_sort_two(arr []int) []int {
  for i := 1; i < len(arr); i++ {
    j := i
    for j > 0 && arr[j - 1] > arr[j] {
      swap(j-1, j, &arr)
      j--;
    }
  }
  return arr

}

func swap(previous int, current int, arr *[]int) {
  arr_tmp := *arr
  copy_arr_tmp := arr_tmp[current]
  arr_tmp[current] = arr_tmp[previous]
  arr_tmp[previous] = copy_arr_tmp
}
