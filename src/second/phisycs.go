package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, vO, sO float64) func(float64) float64{
	time_func := func(t float64) float64{
		return (0.5)*a*math.Pow(t, 2) + vO*t + sO
	}
	return time_func
}


func main(){
	var a, b, c, d float64
    fmt.Printf("Enter value for acceleration: ")
    fmt.Scanf("%f",&a)
    fmt.Printf("Enter value for intial velocity: ")
    fmt.Scanf("%f",&b)
    fmt.Printf("Enter value for intial displacement: ")
    fmt.Scanf("%f",&c)
    fmt.Printf("Enter the value of time: ")
    fmt.Scanf("%f",&d)

    fn := GetDisplaceFn(a,b,c)
    
    fmt.Printf("The displacement is: ")
    fmt.Println(fn(d))
}