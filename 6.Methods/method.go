package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func main() {

	object := Point{X:5, Y:10}
	obj := Point{X:10, Y:20}
	object.Distance(obj)
	fmt.Println("Result : ", object)
}