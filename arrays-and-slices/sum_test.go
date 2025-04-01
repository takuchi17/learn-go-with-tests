package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 5})
		want := 11

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
