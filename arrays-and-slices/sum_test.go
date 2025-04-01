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

func TestSumAll(t *testing.T) {
	t.Run("collections of any size", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3, 5}, []int{2, 3})
		want := []int{11, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
