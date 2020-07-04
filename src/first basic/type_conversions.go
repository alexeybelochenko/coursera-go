package main

import "fmt"

func main() {
	var x int32 = 1
	var y int16 = 2
	x = y // Error. because int32 and int16 its two different types
	x = int32(y) //That`s work

	// float32 ~6 digits of precision
	// float 64 ~15 digits of precision

	//Complex numbers
	var z complex128 = complex(2,3)

	//in go
	//Code poitns - Unicode characters
	//Rune - a code point in Go


	//Iota Example
	type Grades int
	const (
		A Grades= Iota
		B
		C
		D
	)
}