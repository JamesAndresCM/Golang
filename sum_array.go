package main
import "fmt"
func main(){
  nums := []int64{1,2,3,4}
  fmt.Println(AddNums(nums))
}

func AddNums(numbers []int64) int64 {
  var sum int64
  for _,n := range numbers {
    sum+= int64(n)
  }
  return sum
}
