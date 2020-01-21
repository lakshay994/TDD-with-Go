package main

import "math"

// Shape is the common interface to represent the area of different geometrical shapes
type Shape interface {
	Area() float64
}

// Rectangle represents the geometrical shape rectangle
type Rectangle struct {
	length  float64
	breadth float64
}

// Circle represents the geometrical shape circle
type Circle struct {
	radius float64
}

// Perimeter find the perimeter of a given rectangle
func (r Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.length + r.breadth)
}

// Area find the area of a given rectangle
func (r Rectangle) Area() (area float64) {
	return r.length * r.breadth
}

// Area find the area of a given rectangle
func (c Circle) Area() (are float64) {
	return math.Pi * c.radius * c.radius
}
