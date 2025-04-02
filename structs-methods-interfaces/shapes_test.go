package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 36},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 6.0}
		got := Area(rectangle)
		want := 60.0
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
