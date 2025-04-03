package main

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{"Rectangle", Rectangle{12, 6}, 36},
		{"Circle", Circle{10}, 62.83185307179586},
		{"Triangle", Triangle{3, 4, 5}, 12},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("got %g want %g", got, tt.hasPerimeter)
			}
		})
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
