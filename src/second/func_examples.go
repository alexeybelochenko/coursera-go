package main 

import "fmt"


// func foo (x int, y int) {
// 	fmt.Println(x * y)
// }

// func inc(x int) int{
// 	return x + 1
// }

//Multiple Returns 
// func mult(x int) (int, int){
// 	return x, x + 1
// }

//Call by Reference 
// func foo1(y *int ) {
// 	*y++
// }

// func pass_array(x [3]int) int{
	// return x[0]
// }

func array_pointer(x *[3]int){
	(*x)[0] = (*x)[0] + 1
}


func main(){
	a := [3]int{1, 2, 3}
	array_pointer(&a)
	fmt.Println(a)
	// fmt.Println(pass_array(a))
	// foo(2, 3)
	// y := inc(5)
	// fmt.Println(y)
	// a, b := mult(5)
	// x := 2
	// foo1(&x)
	// fmt.Println(x)
}