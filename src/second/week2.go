package main

import "fmt"

// Functions as Arguments
func applyIt(afunc func (int) int, val int) int{
	return afunc(val)
}

func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}


func main(){
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 2))
	//Lambda function
	v := applyIt(func (x int) int {return x + 1}, 2)
	fmt.Println(v)
}