package main

import "fmt"

func main(){
  arr := [10]int{1,234,12,12,3,4,12,456,1,123}
  total := 0  
  for _, value := range arr{
    total += value
  }
  fmt.Println(total)
  append_slice()
  copy_slice()
}

func append_slice(){
  slice_one := make([]int,3)
  slice_one[0] = 2
  slice_one[1] = 3
  slice_one[2] = 4

  slice_two := make([]int,2)
  slice_two[0] = 10
  slice_two[1] = 20

  fmt.Println(slice_one)
  fmt.Println(slice_two)

  slice_three := append(slice_one, 2,12)
  fmt.Println("Slice 3: ",slice_three)

  for _,val := range slice_two {
    slice_one = append(slice_one,val)
  }
  fmt.Println(slice_one)

  var vals []int
  for i := 0; i < 5; i++ {
    vals = append(vals, i)
    fmt.Println(len(vals))
    fmt.Println(cap(vals))
  }
  fmt.Println(vals)
}

func copy_slice(){
  slice_one := []int{1,2,33,1}
  slice_two := make([]int,2)

  copy(slice_two, slice_one)
  fmt.Println(slice_one,slice_two)

}
