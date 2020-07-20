package main

import "fmt"


type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}
type Rectangle {...}
func (t Rectangle) Area() float64 {...}
func (t Rectangle) Perimeter() float64 {...}

func FitInYard(s Shape2D) bool {
	if (s.Area() > 100 && s.Perimeter > 100) {
		return true
	}
	return false
}

// func DrawShape(s Shape2D) bool {
	// rect, ok := s.(Rectangle)
	// if ok {
		// DrawRect(rect)
	// }
	// tri, ok := s.(Triangle)
	// if ok {
		// DrawTri(tri)
	// }
// }
 

func DrawShape(s Shape2D) bool {
	switch := sh := s.(type) {
	case Rectangle:
		DrawRect(sh)
	case Triangle:
		DrawRect(sh)
	}
}



type Speaker interface {Speak()}

// type Dog struct {name string}
// func (d Dog) Speak() {
// 	fmt.Println(d.name)
// }

// func main() {
// 	var s1 Speaker
// 	var d1 Dog("Brian")
// 	s1 = d1
// 	s1.Speak()
// }