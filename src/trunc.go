package main

import "fmt"

func main(){
	var input float64
	fmt.Printf("Enter a float number: ")
	_, err := fmt.Scan(&input)
	answer := int(input)

	if err != nil {
	   fmt.Println(err)
	}
	fmt.Println(answer)
}