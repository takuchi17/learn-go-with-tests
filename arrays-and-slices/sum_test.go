package main

import (
	"reflect"
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

func TestSumAllTails(t *testing.T) {
	t.Run("make the sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3, 5}, []int{2, 3})
		want := []int{10, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("safety sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3})
		want := []int{0, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
