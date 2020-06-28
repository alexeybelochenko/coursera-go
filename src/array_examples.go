package main

import "fmt"

func main() {
	var x[5] int
	y := [3]int {1, 2, 3}
	
	for index, value range y{
		fmt.Println("%d %d", index, value)
	} 

}