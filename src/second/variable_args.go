package main

import "fmt"

func getMax(vals ... int) int {
	maxV := -1
	for _, v := range vals{
		if v > maxV { maxV = v}
	}
	return maxV
}


func main() {
	defer fmt.Println("Bye!") //Execute when main func actually completes
	fmt.Println(getMax(10, 1, 2, 3))
	vslice := []int{1, 3, 4, 100}
	fmt.Println(getMax(vslice...))


}