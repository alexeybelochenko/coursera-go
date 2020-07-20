package main 

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p Point) DistToOrig() float64 {
	t := math.Pow(p.x, 2) + math.Pow(p.y, 2)
	return math.Sqrt(t)
}


func main() {
	r := Point{x: 10, y: 5}
	fmt.Println(r.DistToOrig())
}