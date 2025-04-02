package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("Cicle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
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
