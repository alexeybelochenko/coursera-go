package main

import (
	"fmt"
)

func main() {
	//var x[5] int
	// y := [3]int{1, 2, 3}
	
	// for index, value range y{
	// 	fmt.Println("%d %d", index, value)
	// } 

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	s1 := arr[1:3]
	//s2 := arr[2:5]
	fmt.Println(s1)

}