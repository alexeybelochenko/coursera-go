package main

import "fmt"

func main() {
	for i:=0; i<10; i++{
		fmt.Pritnln("hi ")
	}
	i = 0
	for i < 10 {
		fmt.Pritnln("hi")
		i++
	}
	for { //Infinite Loop
		fmt.Pritnln("hi")
	}

	switch {
	case x > 1:
		fmt.Pritnln("case 1")
	case x < -1:
		fmt.Pritnln("case 2")
	}
}