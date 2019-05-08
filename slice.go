package main

import "fmt"

func main(){
  slice_one := []int{1,2,3,4}
  arreglo := [3]int{1,2,3}

  slice := arreglo[1:]
  fmt.Println(slice_one)
  fmt.Println(slice)

  slice_two := make([]int,3,5)
  // puntero
  // largo
  // capacity capacidad de slice

  slice = append(slice_two,2)
  fmt.Println(cap(slice_two))
  fmt.Println(len(slice_two))
  fmt.Println(slice)
  copy_arrays()
}

func copy_arrays(){
  arr_one := []int{1,2,3,4,5,6}
  arr_two := make([]int, len(arr_one), cap(arr_one)*2)
  //destino, origen
  copy(arr_two,arr_one)

  fmt.Println(arr_one)
  fmt.Println(arr_two)
}
