package main

import "math"

type Shape interface {
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	radias float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radias
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
