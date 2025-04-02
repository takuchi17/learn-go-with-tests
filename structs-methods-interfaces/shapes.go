package main

import "math"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radias float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radias
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
