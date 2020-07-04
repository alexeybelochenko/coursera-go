package main

import "fmt"

func BubbleSort(nums []int) []int {
   for i := len(nums); i > 0; i-- {
      for j := 1; j < i; j++ {
         if nums[j-1] > nums[j] {
            temp := nums[j]
            nums[j] = nums[j-1]
            nums[j-1] = temp
                        }
            }
      	}
   	return nums
}
func main() {
    a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
    fmt.Println(BubbleSort(a))
}