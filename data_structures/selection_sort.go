/*
  0(n2)
  identificar minimo e ir intercambiando porcion ordenada y desordenada
  [123,34,12,2,334]
  min = 2
  min intercambio con porcion desordenada 123
  [2,123,34,12,334]
  min = 12
  min intercambio con 123
  [2,12,123,34,334]
  min = 34
  intercambio con porcion desordenada 123
  [2,12,34,123,334]
  min = 123
*/
package main
import "fmt"

func main(){
  original_arr := []int{412,12,123,1,2324,3,155,23412,2}
  fmt.Println(original_arr)
  order_arr := selection_sort(original_arr)
  fmt.Println(order_arr)
}

func selection_sort(arr []int) []int {
  for i := 0; i < len(arr); i++ {
    min, min_position := arr[i], i
    original_val := arr[i]
    for j := i + 1; j < len(arr); j++{
      comp := arr[j]
      if comp < min {
        min, min_position = comp, j
      }
    }

    if original_val != min {
      arr[i], arr[min_position] = min, original_val
    }
  }
  return arr
}
