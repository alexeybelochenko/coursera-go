/*
Race condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, add_one and sub_one are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/

package main 

import "fmt"

func add(value *int){
	(*value)++
	fmt.Println(*value)
}

func sub(value *int){
	(*value)--
	fmt.Println(*value)
}

func main(){
	val := 0
	for i:= 0; i < 5; i++{
		go add(&val)
		go sub(&val)
		val++
	} 
}